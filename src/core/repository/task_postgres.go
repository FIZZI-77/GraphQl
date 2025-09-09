package repository

import (
	modelsService "GraphQL/src/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

type TaskPostgres struct {
	db *sql.DB
}

func NewTaskPostgres(db *sql.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (t *TaskPostgres) CreateTask(ctx context.Context, userID uuid.UUID, title string, description string) (*modelsService.Tasks, error) {
	task := &modelsService.Tasks{
		UserID:      userID,
		Title:       title,
		Description: description,
	}
	const createTask = `INSERT INTO tasks(user_id, title, description) VALUES ($1, $2, $3)`
	_, err := t.db.ExecContext(ctx, createTask, userID, title, description)
	if err != nil {
		return nil, fmt.Errorf("failed create task: %w", err)
	}
	fmt.Println("task created")
	return task, nil
}
func (t *TaskPostgres) MarkTaskCompleted(ctx context.Context, taskID int32) error {
	const markTaskCompleted = `UPDATE tasks SET completed = true WHERE id = $1`
	_, err := t.db.ExecContext(ctx, markTaskCompleted, taskID)
	if err != nil {
		return fmt.Errorf("failed mark task completed: %w", err)
	}
	fmt.Println("task completed")
	return nil
}
func (t *TaskPostgres) GetTasksByUser(ctx context.Context, userID uuid.UUID) ([]*modelsService.Tasks, error) {
	var userTasks []*modelsService.Tasks

	const getTasksByUser = `SELECT * FROM tasks WHERE user_id = $1`
	rows, err := t.db.QueryContext(ctx, getTasksByUser, userID)
	if err != nil {
		return nil, fmt.Errorf("failed get tasks by user: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		task := &modelsService.Tasks{}
		erro := rows.Scan(&task.UserID, &task.Title, &task.Description)
		if erro != nil {
			return nil, fmt.Errorf("failed get all tasks: %w", err)
		}
		userTasks = append(userTasks, task)
	}
	return userTasks, nil
}
