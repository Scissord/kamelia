package client

import (
	"database/sql"
	"fmt"

	types "auth-microservice/internal/schema/client"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) Get(input types.GetClientsInput) (types.GetClientsResponse, error) {
	var resp types.GetClientsResponse

	query := `
		SELECT id, first_name, last_name, email
		FROM app.client
		ORDER BY id DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.DB.Query(query, input.Limit, input.Offset)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var c types.Client
		if err := rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email); err != nil {
			return resp, err
		}
		resp.Items = append(resp.Items, c)
	}

	if err := rows.Err(); err != nil {
		return resp, err
	}

	countQuery := `SELECT COUNT(*) FROM app.client`

	if err := r.DB.QueryRow(countQuery).Scan(&resp.Total); err != nil {
		return resp, err
	}

	return resp, nil
}

func (r *Repository) Create(input *types.ClientCreateInput) (*types.Client, error) {
	query := `
		INSERT INTO app.client (first_name, last_name, email)
		VALUES ($1, $2, $3)
		RETURNING id, first_name, last_name, email
	`

	var c types.Client
	err := r.DB.QueryRow(query,
		input.FirstName,
		input.LastName,
		input.Email,
	).Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *Repository) Update(input *types.Client) (*types.Client, error) {
	query := `
		UPDATE app.client
		SET first_name = $1,
			last_name = $2,
			email = $3
		WHERE id = $4
		RETURNING id, first_name, last_name, email
	`

	var c types.Client
	err := r.DB.QueryRow(query,
		input.FirstName,
		input.LastName,
		input.Email,
		input.ID,
	).Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *Repository) Delete(id int64) error {
	query := `
		DELETE FROM app.client WHERE id = $1
	`

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("client not found")
	}

	return nil
}
