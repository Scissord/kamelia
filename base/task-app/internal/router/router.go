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

	clientRepo := repository.NewClientRepository(db)
	clientHand := handlers.NewClientHandler(clientRepo)

	// categoryRepo := repository.NewCategoryRepository(db)
	// categoryHand := handlers.NewCategoryHandler(categoryRepo)

	// announcementRepo := repository.NewAnnouncementRepository(db)
	// announcementHand := handlers.AnnouncementHandler(announcementRepo)

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

	// 1. MAIN PAGE
	// GET CATEGORIES
	// GET ANNOUNCEMENT - ?sort=popular

	// r.Route("/categories", func(r chi.Router) {
	// 	r.Get("/", categoryHand.GetCategories)
	// })

	// r.Route("/announcements", func(r chi.Router) {
	// 	r.Get("/", announcementHand.GetAnnouncement)
	// })

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", taskHand.GetTasks)
		r.Post("/", taskHand.CreateTask)

		r.Get("/{id}", taskHand.GetTaskById)
		r.Delete("/{id}", taskHand.DeleteTask)
		r.Patch("/{id}", taskHand.UpdateTask)
	})

	r.Route("clients", func(r chi.Router) {
		r.Get("/", clientHand.GetClients)
		r.Post("/", clientHand.CreateClient)

		r.Get("/{id}", clientHand.GetClientById)
		r.Delete("/{id}", clientHand.DeleteClient)
		r.Patch("/{id}", clientHand.UpdateClient)
	})

	return r
}
