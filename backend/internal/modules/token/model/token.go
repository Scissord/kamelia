package model

import (
	"time"
)

type Token struct {
	ID               uint   `gorm:"primaryKey"`
	UserID           int64  `gorm:"not null;index"`
	RefreshTokenHash string `gorm:"size:64;not null;uniqueIndex"`
	CreatedAt        time.Time
	ExpiresAt        time.Time
	RevokedAt        *time.Time
}

func (Token) TableName() string {
	return "auth.token"
}
