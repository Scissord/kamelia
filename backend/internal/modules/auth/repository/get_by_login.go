package repository

import (
	userModel "backend/internal/modules/user/model"
	"context"
)

func (r *repository) GetByLogin(
	ctx context.Context,
	login string,
) (*userModel.User, error) {

	var user userModel.User

	err := r.db.WithContext(ctx).
		Where("login = ?", login).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
