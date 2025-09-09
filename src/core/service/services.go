package service

import (
	"GraphQL/src/core/repository"
	modelsService "GraphQL/src/models"
	"context"
	"github.com/google/uuid"
)

type User interface {
	GetUserByID(ctx context.Context, userID uuid.UUID) (*modelsService.Users, error)
	GetAllUsers(ctx context.Context) ([]*modelsService.Users, error)
	CreateUser(ctx context.Context, name string, email string) (*modelsService.Users, error)
}

type Task interface {
	CreateTask(ctx context.Context, userID uuid.UUID, title string, description string) (*modelsService.Tasks, error)
	MarkTaskCompleted(ctx context.Context, taskID int32) (*modelsService.Tasks, error)
	GetTasksByUser(ctx context.Context, userID uuid.UUID) ([]*modelsService.Tasks, error)
}

type Service struct {
	User
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos),
		Task: NewTaskService(repos),
	}
}
