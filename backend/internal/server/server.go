package server

import (
	"backend/internal/client"
	"backend/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	clientModule := client.NewModule(cfg.DB)

	// Register routes
	r.Route("/api", func(api chi.Router) {
		api.Mount("/clients", clientModule.Routes())
	})

	return &Server{Router: r}
}
