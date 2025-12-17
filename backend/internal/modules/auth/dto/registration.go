package dto

type RegistrationUserDTO struct {
	Login    string  `json:"login" validate:"required,min=2"`
	Password string  `json:"login" validate:"required,min=2"`
	Name     string  `json:"name" validate:"required,min=2"`
	Email    string  `json:"email" validate:"required,email"`
	Phone    *string `json:"phone" validate:"omitempty"`
	Address  *string `json:"address" validate:"omitempty"`
}
