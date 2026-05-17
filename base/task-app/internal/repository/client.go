package repository

import (
	"context"
	"task-app/internal/models"

	"github.com/jackc/pgx/v5"
)

type ClientRepository struct {
	DB *pgx.Conn
}

func NewClientRepository(db *pgx.Conn) *ClientRepository {
	return &ClientRepository{DB: db}
}

// 1) GET CLIENTS
func (r *ClientRepository) GetClients() ([]models.Client, error) {
	// 1. GET ROWS FROM DB
	rows, err := r.DB.Query(context.Background(), `
		SELECT
			id,
			first_name,
			last_name,
			middle_name,
			phone,
			email,
			created_at,
			updated_at
		FROM app.client
		ORDER BY id DESC;
	`)

	// 2. CHECK IF DB THROW ERROR
	if err != nil {
		return nil, err
	}

	// 3. CLOSE CONNECTION WITH DB
	defer rows.Close()

	// 4. CREATE VARIABLE FOR ROWS
	var clients []models.Client

	// 5. FILL VARIABLE WITH ARRAY FROM DB
	for rows.Next() {
		// 6. CREATE VARIABLE FOR 1 ROW
		var client models.Client

		// 7. FILL VARIABLE WITH ROW
		err := rows.Scan(
			&client.ID,
			&client.FirstName,
			&client.LastName,
			&client.MiddleName,
			&client.Phone,
			&client.Email,
			&client.CreatedAt,
			&client.UpdatedAt,
		)

		// 8. CHECK FOR ERROR
		if err != nil {
			return nil, err
		}

		// 9. ADD TO END OF ARRAY
		clients = append(clients, client)
	}

	// 10. CHECK FOR ERROR WHILE MAPPING
	if rows.Err() != nil {
		return nil, err
	}

	// 11. RETURN ARRAY TO HANDLER
	return clients, nil
}

// 2) CREATE CLIENT
func (r *ClientRepository) CreateClient(input models.CreateClientInput) (models.Client, error) {
	// 1. VARIABLES FOR DB FIELDS
	var id int
	var created_at string
	var updated_at string

	// 2. REQUEST TO DB AND FILL VARIABLES
	err := r.DB.QueryRow(
		context.Background(),
		`
			INSERT INTO app.client (
				first_name,
				last_name,
				middle_name,
				phone,
				email
			) VALUES (
				$1, $2, $3, $4, $5
			) RETURNING
				id,
				created_at,
				updated_at;
		`,
		input.FirstName,
		input.LastName,
		input.MiddleName,
		input.Phone,
		input.Email,
	).Scan(&id, &created_at, &updated_at)

	// 3. CHECK FOR ERROR IN DB
	if err != nil {
		return models.Client{}, err
	}

	// 4. RETURN MODEL TO HANDLER
	return models.Client{
		ID:         id,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		MiddleName: input.MiddleName,
		Phone:      input.Phone,
		Email:      input.Email,
		CreatedAt:  created_at,
		UpdatedAt:  updated_at,
	}, nil
}

// 3) GET CLIENT
func (r *ClientRepository) GetClientById(id int) (models.Client, error) {
	// 1. CREATE VARIABLE TO FILL
	var client models.Client

	// 2. REQUEST TO DB AND FILL VARIABLE
	err := r.DB.QueryRow(context.Background(), `
		SELECT
			id,
			first_name,
			last_name,
			middle_name,
			phone,
			email,
			created_at,
			updated_at
		FROM app.client
		WHERE id = $1;
	`, id).Scan(
		&client.ID,
		&client.FirstName,
		&client.LastName,
		&client.MiddleName,
		&client.Phone,
		&client.Email,
		&client.CreatedAt,
		&client.UpdatedAt,
	)

	// 3. CHECK FOR ERROR FROM DB
	if err != nil {
		return models.Client{}, err
	}

	// 4. RETURN TO HANDLER
	return client, nil
}

// 4) DELETE CLIENT
func (r *ClientRepository) DeleteClient(id int) error {
	// 1. REQUEST TO DB
	result, err := r.DB.Exec(context.Background(), `
		DELETE FROM app.client
		WHERE id = $1;
	`, id)

	// 2. CHECK FOR ERROR IN DB
	if err != nil {
		return err
	}

	// 3. CHECK IF WE REALLY AFFECTED ROWS
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	// 4. RETURN SUCCESS
	return nil
}

// 5) UPDATE CLIENT
func (r *ClientRepository) UpdateClient(client models.Client) error {
	// 1. QUERY TO DB TO UPDATE ROW
	result, err := r.DB.Exec(
		context.Background(),
		`
			UPDATE app.client
			SET first_name = $1,
					last_name = $2,
					middle_name = $3,
					phone = $4,
					email = $5
			WHERE id = $6;
		`,
		client.FirstName,
		client.LastName,
		client.MiddleName,
		client.Phone,
		client.Email,
	)

	// 2. CHECK FOR ERROR FROM DB
	if err != nil {
		return nil
	}

	// 3. CHECK IF WE REALLY AFFECTED ROWS
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	// 4. RETURN SUCCESS
	return nil
}
