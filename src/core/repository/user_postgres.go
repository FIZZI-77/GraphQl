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

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) GetUserByID(ctx context.Context, userID uuid.UUID) (*modelsService.Users, error) {
	start := time.Now()
	user := &modelsService.Users{}
	const getUserByID = `SELECT * FROM users WHERE id = $1`
	err := u.db.QueryRowContext(ctx, getUserByID, userID).Scan(&user.UserID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	duration := time.Since(start).Seconds()
	metrics.DBQueryDuration.WithLabelValues("GetUserByID").Observe(duration)

	if err != nil {
		metrics.DBQueriesTotal.WithLabelValues("GetUserByID", "error").Inc()
		return nil, fmt.Errorf("repo: failed get user by id: %w", err)
	}
	metrics.DBQueriesTotal.WithLabelValues("GetUserByID", "success").Inc()
	return user, nil
}
func (u *UserPostgres) GetAllUsers(ctx context.Context) ([]*modelsService.Users, error) {
	start := time.Now()
	var users []*modelsService.Users
	const getAllUsers = `SELECT * FROM users`
	rows, err := u.db.QueryContext(ctx, getAllUsers)
	duration := time.Since(start).Seconds()
	metrics.DBQueryDuration.WithLabelValues("GetAllUsers").Observe(duration)
	if err != nil {
		metrics.DBQueriesTotal.WithLabelValues("GetAllUsers", "error").Inc()
		return nil, fmt.Errorf("repo: failed get all users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := &modelsService.Users{}
		erro := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if erro != nil {
			return nil, fmt.Errorf("repo: failed get all users: %w", erro)
		}
		users = append(users, user)
	}

	metrics.DBQueriesTotal.WithLabelValues("GetAllUsers", "success").Inc()
	return users, nil
}
func (u *UserPostgres) CreateUser(ctx context.Context, userID uuid.UUID, name string, email string) (*modelsService.Users, error) {
	start := time.Now()
	user := &modelsService.Users{
		UserID: userID,
		Name:   name,
		Email:  email,
	}

	const createUser = `INSERT INTO users (id, name, email) VALUES ($1, $2, $3)`
	_, err := u.db.ExecContext(ctx, createUser, userID, name, email)
	duration := time.Since(start).Seconds()
	metrics.DBQueryDuration.WithLabelValues("CreateUser").Observe(duration)
	if err != nil {
		metrics.DBQueriesTotal.WithLabelValues("CreateUser", "error").Inc()
		return nil, fmt.Errorf("repo: failed create user: %w", err)
	}
	fmt.Println("user created")
	metrics.DBQueriesTotal.WithLabelValues("CreateUser", "success").Inc()
	return user, nil
}
