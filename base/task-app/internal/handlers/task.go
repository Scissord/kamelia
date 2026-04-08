package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"task-app/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type Handler struct {
	DB *pgx.Conn
}

func NewHandler(db *pgx.Conn) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query(context.Background(), `
		SELECT
			id,
			title,
			done
		FROM tasks
		ORDER BY id
	`)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to fetch tasks")
		return
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task

		err := rows.Scan(&task.ID, &task.Title, &task.Done)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "failed to scan task")
			return
		}

		tasks = append(tasks, task)
	}

	if rows.Err() != nil {
		writeError(w, http.StatusInternalServerError, "rows iteration error")
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

	var id int

	err = h.DB.QueryRow(
		context.Background(),
		`
			INSERT INTO tasks (
				title,
				done
			) VALUES (
				$1,
				$2
			)
			RETURNING id
		`,
		input.Title,
		input.Done,
	).Scan(&id)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create task")
		return
	}

	task := models.Task{
		ID:    id,
		Title: input.Title,
		Done:  input.Done,
	}

	writeJSON(w, http.StatusCreated, task)
}

func (h *Handler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromChi(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var task models.Task

	err = h.DB.QueryRow(
		context.Background(),
		`
			SELECT
				id,
				title,
				done
			FROM tasks
			WHERE id = $1
		`,
		id,
	).Scan(&task.ID, &task.Title, &task.Done)

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

	result, err := h.DB.Exec(
		context.Background(),
		`
			DELETE FROM tasks
			WHERE id = $1
		`,
		id,
	)

	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to delete task")
		return
	}

	if result.RowsAffected() == 0 {
		writeError(w, http.StatusNotFound, "task not found")
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

	var task models.Task
	err = h.DB.QueryRow(
		context.Background(),
		`
			SELECT
				id,
				title,
				done
			FROM tasks
			WHERE id = $1
		`,
		id,
	).Scan(&task.ID, &task.Title, &task.Done)

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

	_, err = h.DB.Exec(
		context.Background(),
		`
			UPDATE tasks
			SET title = $2,
					done = $3
			WHERE id = $1
		`,
		id,
		task.Title,
		task.Done,
	)

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
