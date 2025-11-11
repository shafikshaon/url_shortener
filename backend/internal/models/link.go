package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Link struct {
	ID             int64      `json:"id" db:"id"`
	UserID         int64      `json:"user_id" db:"user_id"`
	ShortCode      string     `json:"short_code" db:"short_code"`
	DestinationURL string     `json:"destination_url" db:"destination_url"`
	Title          *string    `json:"title,omitempty" db:"title"`
	Tags           StringList `json:"tags" db:"tags"`
	ExpiresAt      *time.Time `json:"expires_at,omitempty" db:"expires_at"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
}

// StringList is a custom type to handle PostgreSQL TEXT[] arrays
type StringList []string

// Scan implements the sql.Scanner interface
func (s *StringList) Scan(src interface{}) error {
	if src == nil {
		*s = []string{}
		return nil
	}

	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, s)
	case string:
		return json.Unmarshal([]byte(v), s)
	}

	return nil
}

// Value implements the driver.Valuer interface
func (s StringList) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "{}", nil
	}
	return json.Marshal(s)
}

// IsExpired checks if the link has expired
func (l *Link) IsExpired() bool {
	if l.ExpiresAt == nil {
		return false
	}
	return time.Now().After(*l.ExpiresAt)
}

// LinkWithStats includes link data along with analytics
type LinkWithStats struct {
	Link
	TotalClicks int `json:"total_clicks"`
	Last30Days  int `json:"last_30_days"`
}
