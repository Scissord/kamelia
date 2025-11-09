package user

import (
	"database/sql"

	userRepo "auth-microservice/internal/repository/user"
	userService "auth-microservice/internal/service/user"

	"github.com/gorilla/mux"
)

type Module struct {
	db *sql.DB
}

func NewModule(db *sql.DB) *Module {
	return &Module{db: db}
}

func (m *Module) RegisterRoutes(r *mux.Router) {
	repo := userRepo.NewRepository(m.db)
	service := userService.NewService(repo)
	handler := NewHandler(service)
	RegisterRoutes(r, handler)
}
