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
	return &modelsService.Tasks{UserID: userID, Title: title, Description: description}, nil
}
func (u *TaskService) MarkTaskCompleted(ctx context.Context, taskID int32) (*modelsService.Tasks, error) {
	return &modelsService.Tasks{TaskID: taskID, Completed: true}, nil
}
func (u *TaskService) GetTasksByUser(ctx context.Context, userID uuid.UUID) ([]*modelsService.Tasks, error) {
	return []*modelsService.Tasks{}, nil
}
