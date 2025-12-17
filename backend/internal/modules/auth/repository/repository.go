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
	Login(ctx context.Context, user *dto.LoginUserDTO) error

	RegistrationTx(
		ctx context.Context,
		user *userModel.User,
		profile *profileModel.Profile,
	) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Login(ctx context.Context, user *dto.LoginUserDTO) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *repository) RegistrationTx(
	ctx context.Context,
	user *userModel.User,
	profile *profileModel.Profile,
) error {

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := r.createUser(ctx, tx, user); err != nil {
			return err
		}

		if profile != nil {
			profile.UserID = user.ID

			if err := r.createProfile(ctx, tx, profile); err != nil {
				return err
			}
		}

		return nil
	})
}
