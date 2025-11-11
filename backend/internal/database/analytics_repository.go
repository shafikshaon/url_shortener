package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/shafikshaon/url_shortener/internal/models"
)

type AnalyticsRepository struct {
	db *sql.DB
}

func NewAnalyticsRepository(db *sql.DB) *AnalyticsRepository {
	return &AnalyticsRepository{db: db}
}

func (r *AnalyticsRepository) CreateClick(click *models.Click) error {
	query := `
		INSERT INTO clicks (link_id, ip_address, country_code, referer, user_agent, device_type)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, clicked_at
	`

	err := r.db.QueryRow(
		query,
		click.LinkID,
		click.IPAddress,
		click.CountryCode,
		click.Referer,
		click.UserAgent,
		click.DeviceType,
	).Scan(&click.ID, &click.ClickedAt)

	if err != nil {
		return fmt.Errorf("error creating click: %w", err)
	}

	// Update daily analytics cache
	go r.UpdateDailyAnalytics(click.LinkID, click.ClickedAt)

	return nil
}

func (r *AnalyticsRepository) UpdateDailyAnalytics(linkID int64, clickedAt time.Time) error {
	date := clickedAt.Format("2006-01-02")

	query := `
		INSERT INTO link_analytics_daily (link_id, date, click_count)
		VALUES ($1, $2, 1)
		ON CONFLICT (link_id, date)
		DO UPDATE SET click_count = link_analytics_daily.click_count + 1
	`

	_, err := r.db.Exec(query, linkID, date)
	if err != nil {
		return fmt.Errorf("error updating daily analytics: %w", err)
	}

	return nil
}

func (r *AnalyticsRepository) GetLinkStats(linkID int64) (*models.ClickStats, error) {
	stats := &models.ClickStats{}

	// Get total clicks
	err := r.db.QueryRow(
		`SELECT COUNT(*) FROM clicks WHERE link_id = $1`,
		linkID,
	).Scan(&stats.TotalClicks)
	if err != nil {
		return nil, fmt.Errorf("error getting total clicks: %w", err)
	}

	// Get last 30 days clicks
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
	err = r.db.QueryRow(
		`SELECT COUNT(*) FROM clicks WHERE link_id = $1 AND clicked_at >= $2`,
		linkID, thirtyDaysAgo,
	).Scan(&stats.Last30Days)
	if err != nil {
		return nil, fmt.Errorf("error getting last 30 days clicks: %w", err)
	}

	// Get daily clicks (last 30 days)
	stats.DailyClicks, err = r.GetDailyClicks(linkID, 30)
	if err != nil {
		return nil, err
	}

	// Get country stats
	stats.Countries, err = r.GetCountryStats(linkID)
	if err != nil {
		return nil, err
	}

	// Get referer stats
	stats.Referers, err = r.GetRefererStats(linkID)
	if err != nil {
		return nil, err
	}

	// Get device type stats
	stats.DeviceTypes, err = r.GetDeviceTypeStats(linkID)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

func (r *AnalyticsRepository) GetDailyClicks(linkID int64, days int) ([]models.DailyClickCount, error) {
	startDate := time.Now().AddDate(0, 0, -days)

	query := `
		SELECT date, click_count
		FROM link_analytics_daily
		WHERE link_id = $1 AND date >= $2
		ORDER BY date ASC
	`

	rows, err := r.db.Query(query, linkID, startDate.Format("2006-01-02"))
	if err != nil {
		return nil, fmt.Errorf("error getting daily clicks: %w", err)
	}
	defer rows.Close()

	dailyClicks := []models.DailyClickCount{}
	for rows.Next() {
		var dc models.DailyClickCount
		var date time.Time
		err := rows.Scan(&date, &dc.Count)
		if err != nil {
			return nil, fmt.Errorf("error scanning daily click: %w", err)
		}
		dc.Date = date.Format("2006-01-02")
		dailyClicks = append(dailyClicks, dc)
	}

	return dailyClicks, nil
}

func (r *AnalyticsRepository) GetCountryStats(linkID int64) ([]models.CountryStats, error) {
	query := `
		SELECT COALESCE(country_code, 'Unknown') as country_code, COUNT(*) as count
		FROM clicks
		WHERE link_id = $1
		GROUP BY country_code
		ORDER BY count DESC
		LIMIT 10
	`

	rows, err := r.db.Query(query, linkID)
	if err != nil {
		return nil, fmt.Errorf("error getting country stats: %w", err)
	}
	defer rows.Close()

	countries := []models.CountryStats{}
	for rows.Next() {
		var cs models.CountryStats
		err := rows.Scan(&cs.CountryCode, &cs.Count)
		if err != nil {
			return nil, fmt.Errorf("error scanning country stat: %w", err)
		}
		countries = append(countries, cs)
	}

	return countries, nil
}

func (r *AnalyticsRepository) GetRefererStats(linkID int64) ([]models.RefererStats, error) {
	query := `
		SELECT COALESCE(referer, 'Direct') as referer, COUNT(*) as count
		FROM clicks
		WHERE link_id = $1
		GROUP BY referer
		ORDER BY count DESC
		LIMIT 10
	`

	rows, err := r.db.Query(query, linkID)
	if err != nil {
		return nil, fmt.Errorf("error getting referer stats: %w", err)
	}
	defer rows.Close()

	referers := []models.RefererStats{}
	for rows.Next() {
		var rs models.RefererStats
		err := rows.Scan(&rs.Referer, &rs.Count)
		if err != nil {
			return nil, fmt.Errorf("error scanning referer stat: %w", err)
		}
		referers = append(referers, rs)
	}

	return referers, nil
}

func (r *AnalyticsRepository) GetDeviceTypeStats(linkID int64) ([]models.DeviceTypeStats, error) {
	query := `
		SELECT COALESCE(device_type, 'Unknown') as device_type, COUNT(*) as count
		FROM clicks
		WHERE link_id = $1
		GROUP BY device_type
		ORDER BY count DESC
	`

	rows, err := r.db.Query(query, linkID)
	if err != nil {
		return nil, fmt.Errorf("error getting device type stats: %w", err)
	}
	defer rows.Close()

	deviceTypes := []models.DeviceTypeStats{}
	for rows.Next() {
		var dts models.DeviceTypeStats
		err := rows.Scan(&dts.DeviceType, &dts.Count)
		if err != nil {
			return nil, fmt.Errorf("error scanning device type stat: %w", err)
		}
		deviceTypes = append(deviceTypes, dts)
	}

	return deviceTypes, nil
}

// GetUserAnalytics returns analytics for all links owned by a user
func (r *AnalyticsRepository) GetUserAnalytics(userID int64) (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// Total links
	var totalLinks int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM links WHERE user_id = $1`, userID).Scan(&totalLinks)
	if err != nil {
		return nil, err
	}
	stats["total_links"] = totalLinks

	// Total clicks
	var totalClicks int64
	err = r.db.QueryRow(`
		SELECT COUNT(*)
		FROM clicks c
		INNER JOIN links l ON c.link_id = l.id
		WHERE l.user_id = $1
	`, userID).Scan(&totalClicks)
	if err != nil {
		return nil, err
	}
	stats["total_clicks"] = totalClicks

	// Clicks this month
	firstOfMonth := time.Now().Format("2006-01-01")
	var monthClicks int64
	err = r.db.QueryRow(`
		SELECT COUNT(*)
		FROM clicks c
		INNER JOIN links l ON c.link_id = l.id
		WHERE l.user_id = $1 AND c.clicked_at >= $2
	`, userID, firstOfMonth).Scan(&monthClicks)
	if err != nil {
		return nil, err
	}
	stats["clicks_this_month"] = monthClicks

	return stats, nil
}
