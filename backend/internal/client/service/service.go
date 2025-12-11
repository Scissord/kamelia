package service

import (
	"context"

	"backend/internal/client/dto"
	"backend/internal/client/model"
	"backend/internal/client/repository"
	"backend/internal/utils"
)

type Service interface {
	List(ctx context.Context, q dto.GetClientsQuery) ([]model.Client, error)
	GetByID(ctx context.Context, id uint) (*model.Client, error)
	Create(ctx context.Context, c *dto.CreateClientDTO) (*model.Client, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) List(ctx context.Context, q dto.GetClientsQuery) ([]model.Client, error) {
	if q.Page < 1 {
		q.Page = 1
	}

	if q.Limit < 1 {
		q.Limit = 20
	}

	offset := (q.Page - 1) * q.Limit

	allowedSorts := map[string]string{
		"id asc":          "id ASC",
		"id desc":         "id DESC",
		"created_at asc":  "created_at ASC",
		"created_at desc": "created_at DESC",
	}

	sort := "created_at DESC"

	if q.Sort != nil {
		normalized := utils.NormalizeSort(*q.Sort)
		if v, ok := allowedSorts[normalized]; ok {
			sort = v
		}
	}

	return s.repo.List(ctx, q.Limit, offset, sort)
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
