package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const PORT = ":8080"

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks = []Task{
	{ID: 1, Title: "Learn Go", Done: false},
	{ID: 2, Title: "Write Rest Api", Done: false},
	{ID: 3, Title: "Practice JSON", Done: true},
}

func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTasks(w)
		case http.MethodPost:
			createTask(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		id, err := getIDFromPath(r.URL.Path)
		if err != nil {
			http.Error(w, "invalid task id", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			getTaskById(w, id)
		case http.MethodPatch:
			updateTask(w, r, id)
		case http.MethodDelete:
			deleteTask(w, id)
		}
	})

	fmt.Println("Server started on port 8080")
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func getTasks(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTask Task

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	fmt.Printf("Parsed struct: %+v\n", newTask)

	if newTask.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(newTask)
	if err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}

func getTaskById(w http.ResponseWriter, id int) {
	w.Header().Set("Content-Type", "application/json")

	for _, task := range tasks {
		if task.ID == id {
			err := json.NewEncoder(w).Encode(task)
			if err != nil {
				http.Error(w, "failed to encode json", http.StatusInternalServerError)
			}
			return
		}
	}

	http.Error(w, "task not found", http.StatusNotFound)
}

func updateTask(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")

}

func deleteTask(w http.ResponseWriter, id int) {

}

func getIDFromPath(path string) (int, error) {
	idStr := strings.TrimPrefix(path, "/tasks/")
	return strconv.Atoi(idStr)
}
