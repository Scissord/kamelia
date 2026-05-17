package models

type Client struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type CreateClientInput struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

type UpdateClientInput struct {
	FirstName  *string `json:"first_name"`
	LastName   *string `json:"last_name"`
	MiddleName *string `json:"middle_name"`
	Phone      *string `json:"phone"`
	Email      *string `json:"email"`
}
