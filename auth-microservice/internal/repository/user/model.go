package user

import "time"

type User struct {
	ID           int64      `db:"id" json:"id"`
	Login        string     `db:"login" json:"login"`
	PasswordHash string     `db:"password_hash" json:"-"`
	IsActive     bool       `db:"is_active" json:"is_active"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}
