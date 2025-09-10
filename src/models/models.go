package modelsService

import (
	"github.com/google/uuid"
	"time"
)

type Users struct {
	UserID    uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tasks struct {
	TaskID      int
	UserID      uuid.UUID
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
