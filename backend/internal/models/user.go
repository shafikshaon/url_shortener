package models

import (
	"time"

	"gorm.io/gorm"
)

type SubscriptionTier string

const (
	TierFree     SubscriptionTier = "free"
	TierPro      SubscriptionTier = "pro"
	TierBusiness SubscriptionTier = "business"
)

type User struct {
	ID               int64            `json:"id" db:"id" gorm:"primaryKey;autoIncrement"`
	Email            string           `json:"email" db:"email" gorm:"uniqueIndex;not null;size:255"`
	PasswordHash     string           `json:"-" db:"password_hash" gorm:"not null;size:255"`
	APIKey           *string          `json:"api_key,omitempty" db:"api_key" gorm:"uniqueIndex;size:255"`
	SubscriptionTier SubscriptionTier `json:"subscription_tier" db:"subscription_tier" gorm:"type:varchar(50);default:'free'"`
	CreatedAt        time.Time        `json:"created_at" db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time        `json:"updated_at" db:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt   `json:"-" gorm:"index"`
}

// GetLinkLimit returns the maximum number of links allowed for this tier
func (u *User) GetLinkLimit() int {
	switch u.SubscriptionTier {
	case TierFree:
		return 50
	case TierPro:
		return 500
	case TierBusiness:
		return 5000
	default:
		return 50
	}
}

// GetClickLimit returns the monthly click tracking limit for this tier
func (u *User) GetClickLimit() int {
	switch u.SubscriptionTier {
	case TierFree:
		return 1000
	case TierPro:
		return 10000
	case TierBusiness:
		return 100000
	default:
		return 1000
	}
}

// GetAPIRateLimit returns the daily API request limit for this tier
func (u *User) GetAPIRateLimit() int {
	switch u.SubscriptionTier {
	case TierFree:
		return 0 // No API access
	case TierPro:
		return 5000
	case TierBusiness:
		return 50000
	default:
		return 0
	}
}
