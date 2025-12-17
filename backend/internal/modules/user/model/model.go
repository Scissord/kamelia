package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Login        string         `gorm:"column:login;uniqueIndex;not null" json:"login"`
	PasswordHash string         `gorm:"column:password_hash;not null" json:"-"`
	IsActive     bool           `gorm:"column:is_active;default:true" json:"is_active"`
	CreatedAt    time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
}

func (User) TableName() string {
	return "auth.user"
}
