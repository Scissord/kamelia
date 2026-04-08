package models

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type CreateTaskInput struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type UpdateTaskInput struct {
	Title *string `json:"title"`
	Done  *bool   `json:"done"`
}
