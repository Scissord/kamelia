package repository

import (
	"context"

	"task-app/internal/models"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	DB *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(email string, passwordHash string) (models.User, error) {
	var user models.User

	err := r.DB.QueryRow(
		context.Background(),
		`
			INSERT INTO auth.user (
				email, password_hash
			)
			VALUES ($1, $2)
			RETURNING id, email
		`,
		email,
		passwordHash,
	).Scan(&user.ID, &user.Email)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (models.User, string, error) {
	var user models.User
	var passwordHash string

	err := r.DB.QueryRow(
		context.Background(),
		`
			SELECT
				id,
				email,
				password_hash
			FROM auth.user
			WHERE email = $1
		`,
		email,
	).Scan(&user.ID, &user.Email, &passwordHash)

	return user, passwordHash, err
}

func (r *UserRepository) GetUserByID(id int) (models.User, error) {
	var user models.User

	err := r.DB.QueryRow(
		context.Background(),
		`
			SELECT
				id,
				email
			FROM auth.user
			WHERE id = $1
		`,
		id,
	).Scan(&user.ID, &user.Email)

	return user, err
}
