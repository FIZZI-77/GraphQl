package repository

import (
	modelsService "GraphQL/src/models"
	"context"
	"database/sql"
	"github.com/google/uuid"
)

type TaskPostgres struct {
	db *sql.DB
}

func NewTaskPostgres(db *sql.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (t *TaskPostgres) CreateTask(ctx context.Context, userID uuid.UUID, title string, description string) (*modelsService.Tasks, error) {
	return &modelsService.Tasks{}, nil
}
func (t *TaskPostgres) MarkTaskCompleted(ctx context.Context, taskID string) (*modelsService.Tasks, error) {
	return &modelsService.Tasks{}, nil
}
func (t *TaskPostgres) GetTasksByUser(ctx context.Context, userID uuid.UUID) ([]*modelsService.Tasks, error) {
	return []*modelsService.Tasks{}, nil
}
