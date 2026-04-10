package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"task-app/internal/models"
	"task-app/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type Handler struct {
	REPO *repository.TaskRepository
}

func NewHandler(repo *repository.TaskRepository) *Handler {
	return &Handler{REPO: repo}
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.REPO.GetTasks()

	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to fetch tasks")
		return
	}

	writeJSON(w, http.StatusOK, tasks)
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Header.Get("Content-Type") != "application/json" {
		writeError(w, http.StatusUnsupportedMediaType, "content type must be application/json")
		return
	}

	var input models.CreateTaskInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if input.Title == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}

	task, err := h.REPO.CreateTask(input)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create task")
		return
	}

	writeJSON(w, http.StatusCreated, task)
}

func (h *Handler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromChi(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	task, err := h.REPO.GetTaskById(id)

	if err != nil {
		if err == pgx.ErrNoRows {
			writeError(w, http.StatusNotFound, "task not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to get task")
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromChi(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.REPO.DeleteTask(id)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to delete task")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "task deleted",
	})
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, err := getIDFromChi(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var input models.UpdateTaskInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	task, err := h.REPO.GetTaskById(id)

	if err != nil {
		if err == pgx.ErrNoRows {
			writeError(w, http.StatusNotFound, "task not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to get task")
		return
	}

	if input.Title != nil {
		task.Title = *input.Title
	}
	if input.Done != nil {
		task.Done = *input.Done
	}

	err = h.REPO.UpdateTask(task)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to update task")
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func getIDFromChi(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	return strconv.Atoi(idStr)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	if _, err := w.Write(bytes); err != nil {
		fmt.Println("write error:", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{
		"error": message,
	})
}
