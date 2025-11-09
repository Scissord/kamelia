package user

import (
	"errors"

	repo "auth-microservice/internal/repository/user"
)

type Service struct {
	repo *repo.Repository
}

func NewService(r *repo.Repository) *Service {
	return &Service{repo: r}
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

func (s *Service) Create(user *repo.User) error {
	return s.repo.Create(user)
}
