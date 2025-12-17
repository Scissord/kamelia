package dto

type LoginUserDTO struct {
	Login    string `json:"login" validate:"required,min=2"`
	Password string `json:"password" validate:"required,min=2"`
}
