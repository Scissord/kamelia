package repository

import (
	"context"

	userModel "backend/internal/modules/user/model"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	// Login(ctx context.Context, user *dto.LoginUserDTO) error
	Registration(ctx context.Context, user *userModel.User) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// func (r *repository) Login(ctx context.Context, user *dto.LoginUserDTO) error {
// 	return r.db.WithContext(ctx).Create(user).Error
// }

func (r *repository) Registration(ctx context.Context, user *userModel.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
