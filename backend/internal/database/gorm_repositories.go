package database

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/shafikshaon/url_shortener/internal/logger"
	"github.com/shafikshaon/url_shortener/internal/models"
)

// UserRepository implementation using GORM
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	ctx := context.Background()
	logger.Infof(ctx, "Creating user with email: %s", user.Email)

	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		logger.Errorf(ctx, "Failed to create user: %+v", err)
		return fmt.Errorf("error creating user: %w", err)
	}

	logger.Infof(ctx, "Successfully created user with ID: %d", user.ID)
	return nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	ctx := context.Background()
	var user models.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error getting user: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) GetByID(id int64) (*models.User, error) {
	ctx := context.Background()
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error getting user: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) GetByAPIKey(apiKey string) (*models.User, error) {
	ctx := context.Background()
	var user models.User
	if err := r.db.WithContext(ctx).Where("api_key = ?", apiKey).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error getting user: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Model(user).Updates(map[string]interface{}{
		"email":             user.Email,
		"subscription_tier": user.SubscriptionTier,
		"api_key":           user.APIKey,
	}).Error
}

func (r *UserRepository) UpdatePassword(userID int64, passwordHash string) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("password_hash", passwordHash).Error
}

func (r *UserRepository) Delete(id int64) error {
	return r.db.Delete(&models.User{}, id).Error
}

// LinkRepository implementation using GORM
type LinkRepository struct {
	db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) *LinkRepository {
	return &LinkRepository{db: db}
}

func (r *LinkRepository) Create(link *models.Link) error {
	ctx := context.Background()
	logger.Infof(ctx, "Creating link with short code: %s", link.ShortCode)

	if err := r.db.WithContext(ctx).Create(link).Error; err != nil {
		logger.Errorf(ctx, "Failed to create link: %+v", err)
		return fmt.Errorf("error creating link: %w", err)
	}

	logger.Infof(ctx, "Successfully created link with ID: %d", link.ID)
	return nil
}

func (r *LinkRepository) GetByShortCode(shortCode string) (*models.Link, error) {
	ctx := context.Background()
	var link models.Link
	if err := r.db.WithContext(ctx).Where("short_code = ?", shortCode).First(&link).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("link not found")
		}
		return nil, fmt.Errorf("error getting link: %w", err)
	}
	return &link, nil
}

func (r *LinkRepository) GetByID(id int64) (*models.Link, error) {
	var link models.Link
	if err := r.db.First(&link, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("link not found")
		}
		return nil, fmt.Errorf("error getting link: %w", err)
	}
	return &link, nil
}

func (r *LinkRepository) GetByUserID(userID int64, limit, offset int, search string, sortBy string) ([]*models.Link, error) {
	query := r.db.Where("user_id = ?", userID)

	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("short_code ILIKE ? OR destination_url ILIKE ? OR title ILIKE ?",
			searchPattern, searchPattern, searchPattern)
	}

	switch sortBy {
	case "clicks":
		query = query.Select("links.*, (SELECT COUNT(*) FROM clicks WHERE link_id = links.id) as click_count").
			Order("click_count DESC")
	case "created_asc":
		query = query.Order("created_at ASC")
	default:
		query = query.Order("created_at DESC")
	}

	var links []*models.Link
	if err := query.Limit(limit).Offset(offset).Find(&links).Error; err != nil {
		return nil, fmt.Errorf("error getting links: %w", err)
	}
	return links, nil
}

func (r *LinkRepository) CountByUserID(userID int64) (int, error) {
	var count int64
	if err := r.db.Model(&models.Link{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("error counting links: %w", err)
	}
	return int(count), nil
}

func (r *LinkRepository) Update(link *models.Link) error {
	return r.db.Model(link).Updates(map[string]interface{}{
		"destination_url": link.DestinationURL,
		"title":           link.Title,
		"tags":            link.Tags,
		"expires_at":      link.ExpiresAt,
	}).Error
}

func (r *LinkRepository) Delete(id int64, userID int64) error {
	result := r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Link{})
	if result.Error != nil {
		return fmt.Errorf("error deleting link: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("link not found or unauthorized")
	}
	return nil
}

func (r *LinkRepository) ShortCodeExists(shortCode string) (bool, error) {
	var count int64
	if err := r.db.Model(&models.Link{}).Where("short_code = ?", shortCode).Count(&count).Error; err != nil {
		return false, fmt.Errorf("error checking short code: %w", err)
	}
	return count > 0, nil
}

func (r *LinkRepository) GetByTag(userID int64, tag string) ([]*models.Link, error) {
	var links []*models.Link
	if err := r.db.Where("user_id = ? AND ? = ANY(tags)", userID, tag).Order("created_at DESC").Find(&links).Error; err != nil {
		return nil, fmt.Errorf("error getting links by tag: %w", err)
	}
	return links, nil
}

func (r *LinkRepository) GetUserTags(userID int64) ([]string, error) {
	var tags []string
	if err := r.db.Raw("SELECT DISTINCT unnest(tags) as tag FROM links WHERE user_id = ? AND tags IS NOT NULL ORDER BY tag", userID).
		Pluck("tag", &tags).Error; err != nil {
		return nil, fmt.Errorf("error getting user tags: %w", err)
	}

	filteredTags := []string{}
	for _, tag := range tags {
		if tag != "" {
			filteredTags = append(filteredTags, tag)
		}
	}
	return filteredTags, nil
}

// AnalyticsRepository implementation using GORM
type AnalyticsRepository struct {
	db *gorm.DB
}

func NewAnalyticsRepository(db *gorm.DB) *AnalyticsRepository {
	return &AnalyticsRepository{db: db}
}

func (r *AnalyticsRepository) CreateClick(click *models.Click) error {
	ctx := context.Background()
	logger.Infof(ctx, "Creating click for link ID: %d", click.LinkID)

	if err := r.db.WithContext(ctx).Create(click).Error; err != nil {
		logger.Errorf(ctx, "Failed to create click: %+v", err)
		return fmt.Errorf("error creating click: %w", err)
	}

	logger.Infof(ctx, "Successfully created click with ID: %d", click.ID)
	go r.UpdateDailyAnalytics(click.LinkID, click.ClickedAt)
	return nil
}

func (r *AnalyticsRepository) UpdateDailyAnalytics(linkID int64, clickedAt time.Time) error {
	date := clickedAt.Format("2006-01-02")
	return r.db.Exec(`
		INSERT INTO link_analytics_daily (link_id, date, click_count)
		VALUES (?, ?, 1)
		ON CONFLICT (link_id, date)
		DO UPDATE SET click_count = link_analytics_daily.click_count + 1
	`, linkID, date).Error
}

func (r *AnalyticsRepository) GetLinkStats(linkID int64) (*models.ClickStats, error) {
	stats := &models.ClickStats{}

	var totalClicks int64
	r.db.Model(&models.Click{}).Where("link_id = ?", linkID).Count(&totalClicks)
	stats.TotalClicks = totalClicks

	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
	var last30Days int64
	r.db.Model(&models.Click{}).Where("link_id = ? AND clicked_at >= ?", linkID, thirtyDaysAgo).Count(&last30Days)
	stats.Last30Days = last30Days

	var err error
	stats.DailyClicks, err = r.GetDailyClicks(linkID, 30)
	if err != nil {
		return nil, err
	}

	stats.Countries, err = r.GetCountryStats(linkID)
	if err != nil {
		return nil, err
	}

	stats.Referers, err = r.GetRefererStats(linkID)
	if err != nil {
		return nil, err
	}

	stats.DeviceTypes, err = r.GetDeviceTypeStats(linkID)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

func (r *AnalyticsRepository) GetDailyClicks(linkID int64, days int) ([]models.DailyClickCount, error) {
	startDate := time.Now().AddDate(0, 0, -days).Format("2006-01-02")

	var results []struct {
		Date       time.Time
		ClickCount int
	}

	if err := r.db.Table("link_analytics_daily").
		Select("date, click_count").
		Where("link_id = ? AND date >= ?", linkID, startDate).
		Order("date ASC").
		Scan(&results).Error; err != nil {
		return nil, err
	}

	dailyClicks := make([]models.DailyClickCount, len(results))
	for i, result := range results {
		dailyClicks[i] = models.DailyClickCount{
			Date:  result.Date.Format("2006-01-02"),
			Count: result.ClickCount,
		}
	}

	return dailyClicks, nil
}

func (r *AnalyticsRepository) GetCountryStats(linkID int64) ([]models.CountryStats, error) {
	var countries []models.CountryStats
	if err := r.db.Model(&models.Click{}).
		Select("COALESCE(country_code, 'Unknown') as country_code, COUNT(*) as count").
		Where("link_id = ?", linkID).
		Group("country_code").
		Order("count DESC").
		Limit(10).
		Scan(&countries).Error; err != nil {
		return nil, err
	}
	return countries, nil
}

func (r *AnalyticsRepository) GetRefererStats(linkID int64) ([]models.RefererStats, error) {
	var referers []models.RefererStats
	if err := r.db.Model(&models.Click{}).
		Select("COALESCE(referer, 'Direct') as referer, COUNT(*) as count").
		Where("link_id = ?", linkID).
		Group("referer").
		Order("count DESC").
		Limit(10).
		Scan(&referers).Error; err != nil {
		return nil, err
	}
	return referers, nil
}

func (r *AnalyticsRepository) GetDeviceTypeStats(linkID int64) ([]models.DeviceTypeStats, error) {
	var deviceTypes []models.DeviceTypeStats
	if err := r.db.Model(&models.Click{}).
		Select("COALESCE(device_type, 'Unknown') as device_type, COUNT(*) as count").
		Where("link_id = ?", linkID).
		Group("device_type").
		Order("count DESC").
		Scan(&deviceTypes).Error; err != nil {
		return nil, err
	}
	return deviceTypes, nil
}

func (r *AnalyticsRepository) GetUserAnalytics(userID int64) (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	var totalLinks int64
	r.db.Model(&models.Link{}).Where("user_id = ?", userID).Count(&totalLinks)
	stats["total_links"] = totalLinks

	var totalClicks int64
	r.db.Table("clicks c").
		Joins("INNER JOIN links l ON c.link_id = l.id").
		Where("l.user_id = ?", userID).
		Count(&totalClicks)
	stats["total_clicks"] = totalClicks

	firstOfMonth := time.Now().Format("2006-01-01")
	var monthClicks int64
	r.db.Table("clicks c").
		Joins("INNER JOIN links l ON c.link_id = l.id").
		Where("l.user_id = ? AND c.clicked_at >= ?", userID, firstOfMonth).
		Count(&monthClicks)
	stats["clicks_this_month"] = monthClicks

	return stats, nil
}
