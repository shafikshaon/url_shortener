package analytics

import (
	"net/http"
	"strings"

	"github.com/shafikshaon/url_shortener/internal/database"
	"github.com/shafikshaon/url_shortener/internal/models"
)

type Tracker struct {
	analyticsRepo *database.AnalyticsRepository
}

func NewTracker(analyticsRepo *database.AnalyticsRepository) *Tracker {
	return &Tracker{
		analyticsRepo: analyticsRepo,
	}
}

// TrackClick records a click event
func (t *Tracker) TrackClick(linkID int64, ipAddress string, r *http.Request) error {
	click := &models.Click{
		LinkID:    linkID,
		IPAddress: ipAddress,
	}

	// Extract referer
	if referer := r.Header.Get("Referer"); referer != "" {
		click.Referer = &referer
	}

	// Extract user agent
	if userAgent := r.Header.Get("User-Agent"); userAgent != "" {
		click.UserAgent = &userAgent
		deviceType := detectDeviceType(userAgent)
		click.DeviceType = &deviceType
	}

	// Extract country (would require GeoIP database in production)
	// For now, we'll leave it as nil
	// countryCode := getCountryFromIP(ipAddress)
	// click.CountryCode = &countryCode

	return t.analyticsRepo.CreateClick(click)
}

// detectDeviceType determines device type from user agent
func detectDeviceType(userAgent string) string {
	ua := strings.ToLower(userAgent)

	if strings.Contains(ua, "mobile") || strings.Contains(ua, "android") || strings.Contains(ua, "iphone") {
		return "mobile"
	}

	if strings.Contains(ua, "tablet") || strings.Contains(ua, "ipad") {
		return "tablet"
	}

	return "desktop"
}

// GetStats retrieves analytics for a link
func (t *Tracker) GetStats(linkID int64) (*models.ClickStats, error) {
	return t.analyticsRepo.GetLinkStats(linkID)
}

// GetUserAnalytics retrieves overall analytics for a user
func (t *Tracker) GetUserAnalytics(userID int64) (map[string]interface{}, error) {
	return t.analyticsRepo.GetUserAnalytics(userID)
}
