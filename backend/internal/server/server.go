package server

import (
	// "backend/internal/auth"
	"backend/internal/config"
	"backend/internal/middleware"
	"backend/internal/modules/client"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router *chi.Mux
}

func New(cfg *config.Config) *Server {
	r := chi.NewRouter()

	// Глобальные middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60))

	// Init modules
	// authModule := auth.NewModule(cfg.DB)
	clientModule := client.NewModule(cfg.DB)

	// Register routes
	r.Route("/api", func(api chi.Router) {
		// api.Mount("/auth", authModule.Routes())
		// Protected routes
		// api.Use(middleware.AuthMiddleware)
		api.Mount("/clients", clientModule.Routes())
	})

	return &Server{Router: r}
}
