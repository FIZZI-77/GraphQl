package repository

import (
	"GraphQL/metrics"
	modelsService "GraphQL/src/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type TaskPostgres struct {
	db *sql.DB
}

func NewTaskPostgres(db *sql.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (t *TaskPostgres) CreateTask(ctx context.Context, userID uuid.UUID, title string, description string) (*modelsService.Tasks, error) {
	start := time.Now()
	task := &modelsService.Tasks{
		UserID:      userID,
		Title:       title,
		Description: description,
	}
	const createTask = `INSERT INTO tasks(user_id, title, description) VALUES ($1, $2, $3)`
	_, err := t.db.ExecContext(ctx, createTask, userID, title, description)
	duration := time.Since(start).Seconds()

	metrics.DBQueryDuration.WithLabelValues("CreateTask").Observe(duration)

	if err != nil {
		metrics.DBQueriesTotal.WithLabelValues("CreateTask", "error").Inc()
		return nil, fmt.Errorf("repo: failed create task: %w", err)
	}
	fmt.Println("task created")
	metrics.DBQueriesTotal.WithLabelValues("CreateTask", "success").Inc()
	return task, nil
}
func (t *TaskPostgres) MarkTaskCompleted(ctx context.Context, taskID int) (*modelsService.Tasks, error) {
	start := time.Now()
	task := &modelsService.Tasks{
		TaskID: taskID,
	}
	const markTaskCompleted = `UPDATE tasks SET completed = true WHERE id = $1`
	_, err := t.db.ExecContext(ctx, markTaskCompleted, taskID)

	duration := time.Since(start).Seconds()
	metrics.DBQueryDuration.WithLabelValues("MarkTaskCompleted").Observe(duration)
	if err != nil {
		metrics.DBQueriesTotal.WithLabelValues("MarkTaskCompleted", "error").Inc()
		return nil, fmt.Errorf("repo: failed mark task completed: %w", err)
	}
	fmt.Println("task completed")
	metrics.DBQueriesTotal.WithLabelValues("MarkTaskCompleted", "success").Inc()
	return task, nil
}
func (t *TaskPostgres) GetTasksByUser(ctx context.Context, userID uuid.UUID) ([]*modelsService.Tasks, error) {
	start := time.Now()
	var userTasks []*modelsService.Tasks

	const getTasksByUser = `SELECT * FROM tasks WHERE user_id = $1`
	rows, err := t.db.QueryContext(ctx, getTasksByUser, userID)
	duration := time.Since(start).Seconds()
	metrics.DBQueryDuration.WithLabelValues("GetTasksByUser").Observe(duration)
	if err != nil {
		metrics.DBQueriesTotal.WithLabelValues("GetTasksByUser", "error").Inc()
		return nil, fmt.Errorf("repo: failed get tasks by user: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		task := &modelsService.Tasks{}
		erro := rows.Scan(&task.TaskID, &task.UserID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
		if erro != nil {
			return nil, fmt.Errorf("repo: failed get all tasks: %w", err)
		}
		userTasks = append(userTasks, task)
	}
	metrics.DBQueriesTotal.WithLabelValues("GetTasksByUser", "success").Inc()
	return userTasks, nil
}
