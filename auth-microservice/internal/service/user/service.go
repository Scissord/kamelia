package user

import (
	"errors"
	"strings"

	repo "auth-microservice/internal/repository/user"
	types "auth-microservice/internal/schema/auth"
	utils "auth-microservice/internal/utils"
)

type Service struct {
	repo *repo.Repository
}

func NewService(r *repo.Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Registration(input types.RegistrationInput) (*repo.User, error) {
	input.Login = strings.TrimSpace(input.Login)
	input.Password = strings.TrimSpace(input.Password)

	// 1. Validation registration
	if err := s.ValidateRegistration(input); err != nil {
		return nil, err
	}

	// 2. Validate if exist
	existing, err := s.repo.FindByLogin(input.Login)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("user already exists")
	}

	// 3. Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	// 4. Create user
	user := &repo.User{
		Login:        input.Login,
		PasswordHash: string(hashedPassword),
	}

	createdUser, err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *Service) FindByLogin(login string) error {
	existing, err := s.repo.FindByLogin(login)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("user already exists")
	}
	return nil
}
