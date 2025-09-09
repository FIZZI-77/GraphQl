package repository

import (
	modelsService "GraphQL/src/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) GetUserByID(ctx context.Context, userID uuid.UUID) (*modelsService.Users, error) {
	user := &modelsService.Users{}
	const getUserByID = `SELECT * FROM users WHERE id = $1`
	err := u.db.QueryRowContext(ctx, getUserByID, userID).Scan(&user.UserID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed get user by id: %w", err)
	}
	return user, nil
}
func (u *UserPostgres) GetAllUsers(ctx context.Context) ([]*modelsService.Users, error) {
	var users []*modelsService.Users
	const getAllUsers = `SELECT * FROM users`
	rows, err := u.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, fmt.Errorf("failed get all users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := &modelsService.Users{}
		erro := rows.Scan(&user.UserID, &user.Name, &user.Email)
		if erro != nil {
			return nil, fmt.Errorf("failed get all users: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}
func (u *UserPostgres) CreateUser(ctx context.Context, userID uuid.UUID, name string, email string) error {

	const createUser = `INSERT INTO users (id, name, email) VALUES ($1, $2, $3)`
	_, err := u.db.ExecContext(ctx, createUser, userID, name, email)
	if err != nil {
		return fmt.Errorf("failed create user: %w", err)
	}
	fmt.Println("user created")
	return nil
}
