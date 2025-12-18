package repository

import (
	"context"

	"backend/internal/modules/client/model"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	List(ctx context.Context, limit int, offset int, sort string) ([]model.Client, error)
	GetByID(ctx context.Context, id uint) (*model.Client, error)
	Create(ctx context.Context, client *model.Client) error
	Update(ctx context.Context, id uint, client *model.Client) error
	Delete(ctx context.Context, id uint) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) List(ctx context.Context, limit int, offset int, sort string) ([]model.Client, error) {
	var clients []model.Client
	err := r.db.WithContext(ctx).
		Order(sort).
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

func (r *repository) Update(ctx context.Context, id uint, client *model.Client) error {
	return r.db.WithContext(ctx).Model(&model.Client{}).Where("id = ?", id).Updates(client).Error
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Client{}, id).Error
}
