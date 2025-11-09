package user

type RegistrationInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegistrationResponse struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
}
