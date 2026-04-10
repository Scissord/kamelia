package repository

import (
	"context"
	"task-app/internal/models"

	"github.com/jackc/pgx/v5"
)

type TaskRepository struct {
	DB *pgx.Conn
}

func NewTaskRepository(db *pgx.Conn) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) GetTasks() ([]models.Task, error) {
	rows, err := r.DB.Query(context.Background(), `
		SELECT
			id,
			title,
			done
		FROM tasks
		ORDER BY id DESC
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task

		err := rows.Scan(&task.ID, &task.Title, &task.Done)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return tasks, nil
}

func (r *TaskRepository) CreateTask(input models.CreateTaskInput) (models.Task, error) {
	var id int

	err := r.DB.QueryRow(context.Background(), `
		INSERT INTO tasks (
			title, done
		)
		VALUES (
			$1, $2
		)
		RETURNING id;
	`).Scan(&id)

	if err != nil {
		return models.Task{}, err
	}

	return models.Task{
		ID:    id,
		Title: input.Title,
		Done:  input.Done,
	}, nil
}

func (r *TaskRepository) GetTaskById(id int) (models.Task, error) {
	var task models.Task

	err := r.DB.QueryRow(context.Background(),
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
		return models.Task{}, err
	}

	return task, err
}

func (r *TaskRepository) DeleteTask(id int) error {
	result, err := r.DB.Exec(context.Background(),
		`
			DELETE FROM tasks
			WHERE id = $1
		`,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *TaskRepository) UpdateTask(task models.Task) error {
	result, err := r.DB.Exec(
		context.Background(),
		`
			UPDATE tasks
			SET title = $1, done = $2
			WHERE id = $1
		`,
		task.ID,
		task.Title,
		task.Done,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
