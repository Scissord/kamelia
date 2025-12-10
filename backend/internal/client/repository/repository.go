package repository

import (
	"context"

	"backend/internal/client/model"

	"gorm.io/gorm"
)

type Repository interface {
	List(ctx context.Context, limit int, offset int) ([]model.Client, error)
	GetByID(ctx context.Context, id uint) (*model.Client, error)
	Create(ctx context.Context, client *model.Client) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) List(ctx context.Context, limit int, offset int) ([]model.Client, error) {
	var clients []model.Client
	err := r.db.WithContext(ctx).
		Order("id DESC").
		Limit(limit).
		Offset(offset).
		Find(&clients).
		Error

	return clients, err
}

func (r *repository) Create(ctx context.Context, c *model.Client) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *repository) GetByID(ctx context.Context, id uint) (*model.Client, error) {
	var c model.Client
	err := r.db.WithContext(ctx).First(&c, id).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}
