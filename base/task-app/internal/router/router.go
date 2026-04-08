package router

import (
	"net/http"
	"task-app/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func InitRouter(db *pgx.Conn) http.Handler {
	r := chi.NewRouter()

	h := handlers.NewHandler(db)

	r.Get("/tasks", h.GetTasks)
	r.Post("/tasks", h.CreateTask)

	r.Get("/tasks/{id}", h.GetTaskById)
	r.Delete("/tasks/{id}", h.DeleteTask)
	r.Patch("/tasks/{id}", h.UpdateTask)

	return r
}
