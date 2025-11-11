package models

import (
	"time"
)

type Click struct {
	ID          int64     `json:"id" db:"id"`
	LinkID      int64     `json:"link_id" db:"link_id"`
	ClickedAt   time.Time `json:"clicked_at" db:"clicked_at"`
	IPAddress   string    `json:"ip_address" db:"ip_address"`
	CountryCode *string   `json:"country_code,omitempty" db:"country_code"`
	Referer     *string   `json:"referer,omitempty" db:"referer"`
	UserAgent   *string   `json:"user_agent,omitempty" db:"user_agent"`
	DeviceType  *string   `json:"device_type,omitempty" db:"device_type"`
}

type AnalyticsDaily struct {
	LinkID     int64     `json:"link_id" db:"link_id"`
	Date       time.Time `json:"date" db:"date"`
	ClickCount int       `json:"click_count" db:"click_count"`
}

// Analytics response structures
type ClickStats struct {
	TotalClicks   int64                `json:"total_clicks"`
	Last30Days    int64                `json:"last_30_days"`
	DailyClicks   []DailyClickCount    `json:"daily_clicks"`
	Countries     []CountryStats       `json:"countries"`
	Referers      []RefererStats       `json:"referers"`
	DeviceTypes   []DeviceTypeStats    `json:"device_types"`
}

type DailyClickCount struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type CountryStats struct {
	CountryCode string `json:"country_code"`
	Count       int    `json:"count"`
}

type RefererStats struct {
	Referer string `json:"referer"`
	Count   int    `json:"count"`
}

type DeviceTypeStats struct {
	DeviceType string `json:"device_type"`
	Count      int    `json:"count"`
}
