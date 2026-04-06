package service

import (
	userProfileModel "backend/internal/modules/auth/model"
	userModel "backend/internal/modules/user/model"
	"context"
	"net/http"

	"backend/internal/modules/auth/dto"
	"backend/internal/modules/auth/repository"
)

type Service interface {
	Login(ctx context.Context, r *http.Request, dto *dto.LoginUserDTO) (*userModel.User, error)
	Registration(ctx context.Context, dto *dto.RegistrationUserDTO) (*userProfileModel.UserProfile, error)
	// CreateSession(ctx context.Context, dto *dto)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}
