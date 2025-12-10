package client

import (
	"database/sql"

	clientRepo "auth-microservice/internal/repository/client"

	clientService "auth-microservice/internal/service/client"

	"github.com/gorilla/mux"
)

type Module struct {
	db *sql.DB
}

func NewModule(db *sql.DB) *Module {
	return &Module{db: db}
}

func (m *Module) RegisterRoutes(r *mux.Router) {
	// --- Client ---
	cRepo := clientRepo.NewRepository(m.db)
	cService := clientService.NewService(cRepo)

	// --- Handler ---
	handler := NewHandler(cService)

	RegisterRoutes(r, handler)
}
