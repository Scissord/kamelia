package client

import (
	handler "backend/internal/modules/client/handler"
	repository "backend/internal/modules/client/repository"
	service "backend/internal/modules/client/service"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Module struct {
	Handler *handler.Handler
}

func NewModule(db *gorm.DB) *Module {
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

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
