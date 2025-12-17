package model

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Email     string         `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Phone     *string        `gorm:"size:20" json:"phone,omitempty"` // необязательный телефон
	Address   *string        `gorm:"size:512" json:"address,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // для soft delete
}

func (Token) TableName() string {
	return "auth.token"
}
