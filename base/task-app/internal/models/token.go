package models

import "time"

type TokenRecord struct {
	ID        int
	UserID    int
	ExpiresAt time.Time
	RevokedAt *time.Time
}
