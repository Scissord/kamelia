package service

import (
	"context"

	"backend/internal/client/dto"
	"backend/internal/client/model"
	"backend/internal/client/repository"
)

type Service interface {
	List(ctx context.Context, limit int, page int) ([]model.Client, error)
	GetByID(ctx context.Context, id uint) (*model.Client, error)
	Create(ctx context.Context, c *dto.CreateClientDTO) (*model.Client, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) List(ctx context.Context, limit int, page int) ([]model.Client, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 20
	}

	offset := (page - 1) * limit

	return s.repo.List(ctx, limit, offset)
}

func (s *service) GetByID(ctx context.Context, id uint) (*model.Client, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) Create(ctx context.Context, dto *dto.CreateClientDTO) (*model.Client, error) {
	client := model.Client{
		Name:    dto.Name,
		Email:   dto.Email,
		Phone:   dto.Phone,
		Address: dto.Address,
	}

	if err := s.repo.Create(ctx, &client); err != nil {
		return nil, err
	}
	return &client, nil
}
