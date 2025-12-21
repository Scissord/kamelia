package service

import (
	"backend/internal/modules/auth/dto"
	userProfileModel "backend/internal/modules/auth/model"
	profileModel "backend/internal/modules/profile/model"
	userModel "backend/internal/modules/user/model"
	utils "backend/internal/utils"
	"context"
	"errors"
)

var ErrEmailExists = errors.New("Email already exists")
var ErrPhoneExists = errors.New("Phone already exists")

func (s *service) Registration(
	ctx context.Context,
	dto *dto.RegistrationUserDTO,
) (*userProfileModel.UserProfile, error) {
	// Custom validate, check email, phone for unique
	// Check email
	exists, err := s.repo.IsEmailExists(ctx, *dto.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrEmailExists
	}

	// Check phone
	exists, err = s.repo.IsPhoneExists(ctx, *dto.Phone)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrPhoneExists
	}

	// Hash income password
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	// Create model
	user := &userModel.User{
		Login:        dto.Login,
		PasswordHash: hashedPassword,
		IsActive:     true,
	}

	// If any additional info we making profile
	// At least we have Gender, Locale, Timezone
	profile := &profileModel.Profile{
		UserID:     user.ID,
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

	if err := s.repo.RegistrationTx(ctx, user, profile); err != nil {
		return nil, err
	}

	return &userProfileModel.UserProfile{
		ID:         user.ID,
		Login:      user.Login,
		IsActive:   user.IsActive,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		ProfileID:  profile.ID,
		FirstName:  profile.FirstName,
		LastName:   profile.LastName,
		MiddleName: profile.MiddleName,
		Email:      profile.Email,
		Phone:      profile.Phone,
		Birthday:   profile.Birthday,
		Gender:     profile.Gender,
		Locale:     profile.Locale,
		Timezone:   profile.Timezone,
	}, nil
}
