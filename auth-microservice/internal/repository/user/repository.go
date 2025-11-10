package user

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) FindByLogin(login string) (*User, error) {
	var user User

	err := r.DB.QueryRow(`
		SELECT
			id,
			login,
			password_hash
		FROM auth."user"
		WHERE login = $1
	`, login).Scan(&user.ID, &user.Login, &user.PasswordHash)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, err
}

func (r *Repository) Create(user *User) (*User, error) {
	query := `
        INSERT INTO auth."user" (
            login,
            password_hash
        )
        VALUES ($1, $2)
        RETURNING id, login, created_at
    `

	err := r.DB.QueryRow(query, user.Login, user.PasswordHash).Scan(
		&user.ID,
		&user.Login,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
