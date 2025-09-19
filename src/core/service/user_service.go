package service

import (
	"GraphQL/src/core/repository"
	modelsService "GraphQL/src/models"
	"context"
	"github.com/google/uuid"
)

type UserService struct {
	repo repository.CachedRepo
}

func NewUserService(repo repository.CachedRepo) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) GetUserByID(ctx context.Context, userID uuid.UUID) (*modelsService.Users, error) {
	return u.repo.GetUserByID(ctx, userID)
}
func (u *UserService) GetAllUsers(ctx context.Context) ([]*modelsService.Users, error) {
	return u.repo.GetAllUsers(ctx)
}
func (u *UserService) CreateUser(ctx context.Context, name string, email string) (*modelsService.Users, error) {
	id := uuid.New()
	return u.repo.CreateUser(ctx, id, name, email)
}
