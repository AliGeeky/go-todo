package models

import (
	"time"
)

// Task represents a single to-do item.
type Task struct {
	ID          string    `json:"id"`           // Unique identifier for the task
	Title       string    `json:"title"`        // Title or description of the task
	IsCompleted bool      `json:"is_completed"` // Whether the task is completed or not
	CreatedAt   time.Time `json:"created_at"`   // Timestamp when the task was created
}
