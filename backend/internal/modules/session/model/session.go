package model

import (
	"time"
)

type Session struct {
	ID     uint  `gorm:"primaryKey" json:"id"`
	UserID int64 `gorm:"not null;index" json:"user_id"`

	IPAddress string `gorm:"type:inet;not null" json:"ip_address"`
	UserAgent string `gorm:"size:1024" json:"user_agent,omitempty"`

	LoginAt    time.Time  `gorm:"autoCreateTime" json:"login_at"`
	LogoutAt   *time.Time `json:"logout_at,omitempty"`
	LastSeenAt *time.Time `json:"last_seen_at,omitempty"`

	IsActive bool `gorm:"default:true" json:"is_active"`
}

func (Session) TableName() string {
	return "auth.session"
}
