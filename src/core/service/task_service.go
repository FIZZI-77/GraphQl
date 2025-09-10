package service

import (
	"GraphQL/src/core/repository"
	modelsService "GraphQL/src/models"
	"context"
	"github.com/google/uuid"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (u *TaskService) CreateTask(ctx context.Context, userID uuid.UUID, title string, description string) (*modelsService.Tasks, error) {
	return u.repo.CreateTask(ctx, userID, title, description)
}
func (u *TaskService) MarkTaskCompleted(ctx context.Context, taskID int) (*modelsService.Tasks, error) {
	return u.repo.MarkTaskCompleted(ctx, taskID)
}
func (u *TaskService) GetTasksByUser(ctx context.Context, userID uuid.UUID) ([]*modelsService.Tasks, error) {
	return u.repo.GetTasksByUser(ctx, userID)
}
