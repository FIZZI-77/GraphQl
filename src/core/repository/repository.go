package repository

import (
	modelsService "GraphQL/src/models"
	"context"
	"database/sql"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, userID uuid.UUID) (*modelsService.Users, error)
	GetAllUsers(ctx context.Context) ([]*modelsService.Users, error)
	CreateUser(ctx context.Context, userID uuid.UUID, name string, email string) (*modelsService.Users, error)
}

type TaskRepository interface {
	CreateTask(ctx context.Context, userID uuid.UUID, title string, description string) (*modelsService.Tasks, error)
	MarkTaskCompleted(ctx context.Context, taskID int32) error
	GetTasksByUser(ctx context.Context, userID uuid.UUID) ([]*modelsService.Tasks, error)
}

type Repository struct {
	UserRepository
	TaskRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepository: NewUserPostgres(db),
		TaskRepository: NewTaskPostgres(db),
	}
}
