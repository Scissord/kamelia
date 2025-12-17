package repository

import (
	"backend/internal/modules/auth/dto"
	userModel "backend/internal/modules/user/model"
	"backend/internal/utils"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func (r *repository) Login(ctx context.Context, dto *dto.LoginUserDTO) (*userModel.User, error) {
	var user userModel.User

	err := r.db.WithContext(ctx).
		Where("login = ?", dto.Login).
		First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	if !utils.CheckPasswordHash(dto.Password, user.PasswordHash) {
		return nil, fmt.Errorf("invalid password")
	}

	return &user, nil
}
