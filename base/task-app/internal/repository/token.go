package repository

import (
	"context"
	"time"

	"task-app/internal/models"

	"github.com/jackc/pgx/v5"
)

type TokenRepository struct {
	DB *pgx.Conn
}

func NewTokenRepository(db *pgx.Conn) *TokenRepository {
	return &TokenRepository{DB: db}
}

func (r *TokenRepository) CreateRefreshToken(userID int, tokenHash string, expiresAt time.Time) error {
	_, err := r.DB.Exec(
		context.Background(),
		`
			INSERT INTO auth.token (
				user_id,
				token_hash,
				expires_at
			)
			VALUES ($1, $2, $3)
		`,
		userID,
		tokenHash,
		expiresAt,
	)

	return err
}

func (r *TokenRepository) GetByHash(tokenHash string) (models.TokenRecord, error) {
	var token models.TokenRecord

	err := r.DB.QueryRow(
		context.Background(),
		`
			SELECT
				id,
				user_id,
				expires_at,
				revoked_at
			FROM auth.token
			WHERE token_hash = $1
		`,
		tokenHash,
	).Scan(&token.ID, &token.UserID, &token.ExpiresAt, &token.RevokedAt)

	return token, err
}

func (r *TokenRepository) RevokeToken(id int) error {
	_, err := r.DB.Exec(
		context.Background(),
		`
			UPDATE auth.token
			SET revoked_at = NOW()
			WHERE id = $1
		`,
		id,
	)

	return err
}
