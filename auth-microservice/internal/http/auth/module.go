package auth

import (
	"database/sql"

	profileRepo "auth-microservice/internal/repository/profile"
	userRepo "auth-microservice/internal/repository/user"

	profileService "auth-microservice/internal/service/profile"
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
	// repo := userRepo.NewRepository(m.db)
	// service := userService.NewService(repo)
	// handler := NewHandler(service)

	// --- User ---
	uRepo := userRepo.NewRepository(m.db)
	uService := userService.NewService(uRepo)

	// --- Profile ---
	pRepo := profileRepo.NewRepository(m.db)
	pService := profileService.NewService(pRepo)

	// --- Handler ---
	handler := NewHandler(uService, pService)

	RegisterRoutes(r, handler)
}
