package service

import (
	"context"

	// profileModel "backend/internal/modules/profile/model"
	userModel "backend/internal/modules/user/model"
	"backend/internal/utils"

	"backend/internal/modules/auth/dto"
	"backend/internal/modules/auth/repository"
)

type Service interface {
	Registration(ctx context.Context, dto *dto.RegistrationUserDTO) (*userModel.User, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Registration(ctx context.Context, dto *dto.RegistrationUserDTO) (*userModel.User, error) {

	// Тут валидация на дублирование логина, телефона и.т.д
	// Заполняем auth.user
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	user := userModel.User{
		Login:        dto.Login,
		PasswordHash: hashedPassword,
		IsActive:     true,
	}

	if err := s.repo.Registration(ctx, &user); err != nil {
		return nil, err
	}
	return &user, nil

	// Заполняем profile
	// profile := profileModel.Profile{

	// }
}
