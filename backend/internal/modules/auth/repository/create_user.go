package repository

import (
	userModel "backend/internal/modules/user/model"
	"context"

	"gorm.io/gorm"
)

func (r *repository) createUser(
	ctx context.Context,
	tx *gorm.DB,
	user *userModel.User,
) error {
	return tx.WithContext(ctx).Create(user).Error
}
