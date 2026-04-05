package client

import (
	handler "backend/internal/modules/client/handler"
	repository "backend/internal/modules/client/repository"
	service "backend/internal/modules/client/service"
	"fmt"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Module struct {
	Handler *handler.Handler
}

func NewModule(db *gorm.DB) *Module {
	fmt.Printf("db type=%T value=%v\n", db, db)

	repo := repository.NewRepository(db)

	fmt.Printf("repo type=%T value=%v\n", repo, repo)

	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	// Module{Handler: handler} — создаёт значение структуры
	// &Module{...} — возвращает указатель на эту структуру
	return &Module{Handler: handler}
}

func (m *Module) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", m.Handler.List)
	r.Post("/", m.Handler.Create)
	r.Patch("/{id}", m.Handler.Update)
	r.Delete("/{id}", m.Handler.Delete)

	return r
}
