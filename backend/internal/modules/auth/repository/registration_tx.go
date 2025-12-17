package repository

import (
	profileModel "backend/internal/modules/profile/model"
	userModel "backend/internal/modules/user/model"
	"context"

	"gorm.io/gorm"
)

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
