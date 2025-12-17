package service

import (
	"context"

	profileModel "backend/internal/modules/profile/model"
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

func (s *service) Registration(
	ctx context.Context,
	dto *dto.RegistrationUserDTO,
) (*userModel.User, error) {

	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	user := &userModel.User{
		Login:        dto.Login,
		PasswordHash: hashedPassword,
		IsActive:     true,
	}

	var profile *profileModel.Profile

	if dto.FirstName != nil ||
		dto.LastName != nil ||
		dto.MiddleName != nil ||
		dto.Email != nil ||
		dto.Phone != nil ||
		dto.Birthday != nil {

		profile = &profileModel.Profile{
			FirstName:  dto.FirstName,
			LastName:   dto.LastName,
			MiddleName: dto.MiddleName,
			Email:      dto.Email,
			Phone:      dto.Phone,
			Birthday:   dto.Birthday,
			Gender:     dto.Gender,
			Locale:     dto.Locale,
			Timezone:   dto.Timezone,
		}
	}

	if err := s.repo.RegistrationTx(ctx, user, profile); err != nil {
		return nil, err
	}

	return user, nil
}
