package service

import (
	"context"

	userProfileModel "backend/internal/modules/auth/model"

	"backend/internal/modules/auth/dto"
	"backend/internal/modules/auth/repository"
)

type Service interface {
	Registration(ctx context.Context, dto *dto.RegistrationUserDTO) (*userProfileModel.UserProfile, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}
