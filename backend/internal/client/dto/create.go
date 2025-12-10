package dto

type CreateClientDTO struct {
	Name    string  `json:"name" validate:"required,min=2"`
	Email   string  `json:"email" validate:"required,email"`
	Phone   *string `json:"phone" validate:"omitempty"`
	Address *string `json:"address" validate:"omitempty"`
}
