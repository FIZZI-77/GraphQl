package service

import (
	"GraphQL/src/core/repository"
	modelsService "GraphQL/src/models"
	"context"
	"github.com/google/uuid"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) GetUserByID(ctx context.Context, userID uuid.UUID) (*modelsService.Users, error) {
	return &modelsService.Users{UserID: userID}, nil
}
func (u *UserService) GetAllUsers(ctx context.Context) ([]*modelsService.Users, error) {
	return []*modelsService.Users{}, nil
}
func (u *UserService) CreateUser(ctx context.Context, name string, email string) (*modelsService.Users, error) {
	return &modelsService.Users{Name: name, Email: email}, nil
}
