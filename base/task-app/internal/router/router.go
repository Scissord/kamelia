package router

import (
	"net/http"
	"task-app/internal/handlers"
	"task-app/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func InitRouter(db *pgx.Conn) http.Handler {
	r := chi.NewRouter()

	repo := repository.NewTaskRepository(db)
	hand := handlers.NewHandler(repo)

	r.Get("/tasks", hand.GetTasks)
	r.Post("/tasks", hand.CreateTask)

	r.Get("/tasks/{id}", hand.GetTaskById)
	r.Delete("/tasks/{id}", hand.DeleteTask)
	r.Patch("/tasks/{id}", hand.UpdateTask)

	return r
}
