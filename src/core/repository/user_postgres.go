package repository

import (
	modelsService "GraphQL/src/models"
	"context"
	"database/sql"
	"github.com/google/uuid"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) GetUserByID(ctx context.Context, userID uuid.UUID) (*modelsService.Users, error) {
	return &modelsService.Users{}, nil
}
func (u *UserPostgres) GetAllUsers(ctx context.Context) ([]*modelsService.Users, error) {
	return []*modelsService.Users{}, nil
}
func (u *UserPostgres) CreateUser(ctx context.Context, name string, email string) (*modelsService.Users, error) {
	return &modelsService.Users{}, nil
}
