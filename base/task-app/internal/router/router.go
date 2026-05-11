package router

import (
	"net/http"
	"task-app/internal/config"
	"task-app/internal/handlers"
	"task-app/internal/repository"

	authMiddleware "task-app/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

func InitRouter(db *pgx.Conn, cfg *config.Config) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	taskRepo := repository.NewTaskRepository(db)
	taskHand := handlers.NewTaskHandler(taskRepo)

	userRepo := repository.NewUserRepository(db)
	tokenRepo := repository.NewTokenRepository(db)
	authHandler := handlers.NewAuthHandler(userRepo, tokenRepo, cfg)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/registration", authHandler.Register)
		r.Post("/login", authHandler.Login)
		r.Post("/refresh", authHandler.Refresh)
		r.Post("/logout", authHandler.Logout)

		r.Group(func(r chi.Router) {
			r.Use(authMiddleware.AuthMiddleware(cfg))

			r.Get("/me", authHandler.Me)
		})
	})

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", taskHand.GetTasks)
		r.Post("/", taskHand.CreateTask)

		r.Get("/{id}", taskHand.GetTaskById)
		r.Delete("/{id}", taskHand.DeleteTask)
		r.Patch("/{id}", taskHand.UpdateTask)
	})

	return r
}
