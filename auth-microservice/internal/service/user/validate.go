package user

import (
	"errors"
	"strings"

	types "auth-microservice/internal/schema/user"
)

func (s *Service) ValidateRegistration(input types.RegistrationInput) error {
	input.Login = strings.TrimSpace(input.Login)
	input.Password = strings.TrimSpace(input.Password)

	if input.Login == "" || input.Password == "" {
		return errors.New("login and password are required")
	}
	if len(input.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	return nil
}
