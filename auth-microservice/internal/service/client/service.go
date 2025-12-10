package client

import (
	repo "auth-microservice/internal/repository/client"
	types "auth-microservice/internal/schema/client"
)

type Service struct {
	repo *repo.Repository
}

func NewService(r *repo.Repository) *Service {
	return &Service{repo: r}
}

// Получение списка клиентов
func (s *Service) Get(input types.GetClientsInput) (types.GetClientsResponse, error) {
	return s.repo.Get(input)
}

// Создание клиента
func (s *Service) Create(input *types.ClientCreateInput) (*types.Client, error) {
	return s.repo.Create(input)
}

// Обновление клиента
func (s *Service) Update(input *types.Client) (*types.Client, error) {
	return s.repo.Update(input)
}

// Удаление клиента по id
func (s *Service) Delete(id int64) error {
	return s.repo.Delete(id)
}
