package models

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	User        User   `json:"user"`
	AccessToken string `json:"accessToken"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
