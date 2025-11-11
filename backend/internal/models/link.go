package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Link struct {
	ID             int64          `json:"id" db:"id" gorm:"primaryKey;autoIncrement"`
	UserID         int64          `json:"user_id" db:"user_id" gorm:"not null;index"`
	ShortCode      string         `json:"short_code" db:"short_code" gorm:"uniqueIndex;not null;size:255"`
	DestinationURL string         `json:"destination_url" db:"destination_url" gorm:"not null;type:text"`
	Title          *string        `json:"title,omitempty" db:"title" gorm:"size:500"`
	Tags           StringList     `json:"tags" db:"tags" gorm:"type:text[]"`
	ExpiresAt      *time.Time     `json:"expires_at,omitempty" db:"expires_at"`
	CreatedAt      time.Time      `json:"created_at" db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"updated_at" db:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
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
		return s.parseArray(string(v))
	case string:
		return s.parseArray(v)
	default:
		// Fallback to JSON parsing for compatibility
		if jsonBytes, ok := src.([]byte); ok {
			return json.Unmarshal(jsonBytes, s)
		}
	}

	return nil
}

// Value implements the driver.Valuer interface
func (s StringList) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "{}", nil
	}

	// Format as PostgreSQL array: {value1,value2,value3}
	var result string
	result = "{"
	for i, tag := range s {
		if i > 0 {
			result += ","
		}
		// Escape special characters
		escaped := tag
		if containsSpecialChars(tag) {
			escaped = `"` + escapeString(tag) + `"`
		}
		result += escaped
	}
	result += "}"

	return result, nil
}

// parseArray parses PostgreSQL array format
func (s *StringList) parseArray(str string) error {
	if str == "{}" || str == "" {
		*s = []string{}
		return nil
	}

	// Remove braces
	if len(str) > 2 {
		str = str[1 : len(str)-1]
	}

	if str == "" {
		*s = []string{}
		return nil
	}

	// Simple split by comma (can be enhanced for quoted values)
	tags := []string{}
	current := ""
	inQuotes := false

	for i := 0; i < len(str); i++ {
		char := str[i]

		if char == '"' {
			inQuotes = !inQuotes
			continue
		}

		if char == ',' && !inQuotes {
			tags = append(tags, current)
			current = ""
			continue
		}

		current += string(char)
	}

	if current != "" {
		tags = append(tags, current)
	}

	*s = tags
	return nil
}

func containsSpecialChars(s string) bool {
	for _, c := range s {
		if c == ',' || c == '{' || c == '}' || c == '"' || c == ' ' {
			return true
		}
	}
	return false
}

func escapeString(s string) string {
	result := ""
	for _, c := range s {
		if c == '"' || c == '\\' {
			result += "\\"
		}
		result += string(c)
	}
	return result
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
