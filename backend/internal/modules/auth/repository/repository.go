package repository

import (
	"context"

	"backend/internal/modules/auth/dto"
	profileModel "backend/internal/modules/profile/model"
	userModel "backend/internal/modules/user/model"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Login(ctx context.Context, user *dto.LoginUserDTO) (*userModel.User, error)

	IsEmailExists(ctx context.Context, email string) (bool, error)
	IsPhoneExists(ctx context.Context, phone string) (bool, error)

	RegistrationTx(
		ctx context.Context,
		user *userModel.User,
		profile *profileModel.Profile,
	) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
