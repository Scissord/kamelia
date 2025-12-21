package server

import (
	"backend/internal/config"
	"backend/internal/middleware"
	"backend/internal/modules/auth"
	"backend/internal/modules/client"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router *chi.Mux
}

func New(cfg *config.Config) *Server {
	r := chi.NewRouter()
	// Rate limiting – защита от brute force login.
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60))
	r.Use(middleware.CORS)

	authModule := auth.NewModule(cfg.DB)
	clientModule := client.NewModule(cfg.DB)

	r.Route("/api", func(api chi.Router) {
		api.Mount("/auth", authModule.Routes())
		api.Mount("/clients", clientModule.Routes())
	})

	return &Server{Router: r}
}
