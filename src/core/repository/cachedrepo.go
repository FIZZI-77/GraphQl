package repository

import (
	modelsService "GraphQL/src/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"time"
)

type CachedRepo struct {
	userRepo UserRepository
	taskRepo TaskRepository
	redis    *redis.Client
	ttl      time.Duration
}

func NewCachedRepo(dbRepo *Repository, redis *redis.Client, ttl time.Duration) *CachedRepo {
	return &CachedRepo{
		userRepo: dbRepo.UserRepository,
		taskRepo: dbRepo.TaskRepository,
		redis:    redis,
		ttl:      ttl,
	}
}

func (c *CachedRepo) CreateTask(ctx context.Context, userID uuid.UUID, title string, description string) (*modelsService.Tasks, error) {

	key := "task:" + userID.String()
	err := c.redis.Del(ctx, key).Err()
	if err != nil {
		return nil, fmt.Errorf("CachedRepo: CreateTask: failed to delete task: %w", err)
	}
	return c.taskRepo.CreateTask(ctx, userID, title, description)
}

func (c *CachedRepo) MarkTaskCompleted(ctx context.Context, taskID int) (*modelsService.Tasks, error) {
	keys, err := c.redis.Keys(ctx, "tasks:*").Result()
	if err != nil {
		return nil, fmt.Errorf("CachedRepo: MarkTaskCompleted: failed to fetch tasks: %w", err)
	}

	if len(keys) > 0 {
		_, err := c.redis.Del(ctx, keys...).Result()
		if err != nil {
			return nil, fmt.Errorf("CachedRepo: MarkTaskCompleted: failed to delete tasks: %w", err)
		}
	}

	return c.taskRepo.MarkTaskCompleted(ctx, taskID)
}

func (c *CachedRepo) GetTasksByUser(ctx context.Context, userID uuid.UUID) ([]*modelsService.Tasks, error) {
	key := "task:" + userID.String()

	val, err := c.redis.Get(ctx, key).Result()

	if err == nil {
		var tasks []*modelsService.Tasks
		if unMarshalErr := json.Unmarshal([]byte(val), &tasks); unMarshalErr == nil {
			return tasks, nil
		}
	}

	tasks, err := c.taskRepo.GetTasksByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(tasks)
	_ = c.redis.Set(ctx, key, string(data), c.ttl)

	return tasks, nil
}
func (c *CachedRepo) GetUserByID(ctx context.Context, userID uuid.UUID) (*modelsService.Users, error) {
	key := "user:" + userID.String()

	val, err := c.redis.Get(ctx, key).Result()
	if err == nil {
		var users modelsService.Users
		if unMarshalErr := json.Unmarshal([]byte(val), &users); unMarshalErr == nil {
			return &users, nil
		}
	}

	user, err := c.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	data, _ := json.Marshal(user)
	_ = c.redis.Set(ctx, key, string(data), c.ttl)
	return user, nil
}

func (c *CachedRepo) GetAllUsers(ctx context.Context) ([]*modelsService.Users, error) {
	key := "all_users"

	val, err := c.redis.Get(ctx, key).Result()
	if err == nil {
		var users []*modelsService.Users
		if unMarshalErr := json.Unmarshal([]byte(val), &users); unMarshalErr == nil {
			fmt.Println("достали из кеша")
			return users, nil
		}
	}
	users, err := c.userRepo.GetAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("CachedRepo: GetAllUsers: failed to fetch all users: %w", err)
	}
	data, _ := json.Marshal(users)
	_ = c.redis.Set(ctx, key, string(data), c.ttl)
	return users, nil
}

func (c *CachedRepo) CreateUser(ctx context.Context, userID uuid.UUID, name string, email string) (*modelsService.Users, error) {
	key := "all_users"
	_, err := c.redis.Del(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("CachedRepo: CreateUser: failed to delete user: %w", err)
	}
	return c.userRepo.CreateUser(ctx, userID, name, email)
}
