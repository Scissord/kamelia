package auth

import (
	handler "backend/internal/modules/auth/handler"
	repository "backend/internal/modules/auth/repository"
	service "backend/internal/modules/auth/service"

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

	r.Post("/registration", m.Handler.Registration)
	r.Post("/login", m.Handler.Login)

	return r
}
