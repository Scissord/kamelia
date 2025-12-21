package model

import (
	"time"
)

type UserProfile struct {
	// user
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Login     string    `gorm:"column:login;uniqueIndex;not null" json:"login"`
	IsActive  bool      `gorm:"column:is_active;default:true" json:"is_active"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	// profile
	ProfileID  int64   `gorm:"column:profile_id;not null;index" json:"profile_id"`
	FirstName  *string `gorm:"-" json:"first_name,omitempty"`
	LastName   *string `gorm:"-" json:"last_name,omitempty"`
	MiddleName *string `gorm:"-" json:"middle_name,omitempty"`
	Email      *string `gorm:"-" json:"email,omitempty"`
	Phone      *string `gorm:"-" json:"phone,omitempty"`
	Birthday   *string `gorm:"-" json:"birthday,omitempty"`
	Gender     *string `gorm:"column:gender" json:"gender"`
	Locale     *string `gorm:"column:locale" json:"locale"`
	Timezone   *string `gorm:"column:timezone" json:"timezone"`
}
