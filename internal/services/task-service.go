package services

import (
	"errors"
	"time"

	"github.com/AliGeeky/go-todo/internal/models"
	"github.com/AliGeeky/go-todo/internal/repository"
	"github.com/google/uuid" // For generating unique IDs
)

// TaskService defines the business logic for tasks.
type TaskService struct {
	repo repository.TaskRepository
}

// NewTaskService creates a new instance of TaskService.
// It takes a TaskRepository interface as a dependency.
func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// CreateTask creates a new task.
func (s *TaskService) CreateTask(title string) (*models.Task, error) {
	if title == "" {
		return nil, errors.New("task title cannot be empty")
	}

	newTask := models.Task{
		ID:          uuid.New().String(), // Generate a unique ID
		Title:       title,
		IsCompleted: false,
		CreatedAt:   time.Now(),
	}

	err := s.repo.CreateTask(newTask)
	if err != nil {
		return nil, err
	}
	return &newTask, nil
}

// GetTaskByID retrieves a task by its ID.
func (s *TaskService) GetTaskByID(id string) (*models.Task, error) {
	if id == "" {
		return nil, errors.New("task ID cannot be empty")
	}
	return s.repo.GetTaskByID(id)
}

// GetAllTasks retrieves all tasks.
func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

// UpdateTask updates an existing task.
func (s *TaskService) UpdateTask(id, title string, isCompleted bool) (*models.Task, error) {
	if id == "" {
		return nil, errors.New("task ID cannot be empty")
	}
	if title == "" {
		return nil, errors.New("task title cannot be empty")
	}

	existingTask, err := s.repo.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	if existingTask == nil {
		return nil, errors.New("task not found")
	}

	existingTask.Title = title
	existingTask.IsCompleted = isCompleted

	err = s.repo.UpdateTask(*existingTask)
	if err != nil {
		return nil, err
	}
	return existingTask, nil
}

// DeleteTask deletes a task by its ID.
func (s *TaskService) DeleteTask(id string) error {
	if id == "" {
		// Return an error if the task ID is empty.
		// This addresses the "too many return values" error by returning only the 'error' type.
		return errors.New("task ID cannot be empty")
	}

	// Optional: Check if task exists before deleting.
	// Uncomment and implement this block if you want to explicitly check for existence
	// and return a "task not found" error if it doesn't exist.
	// _, err := s.repo.GetTaskByID(id)
	// if err != nil {
	//   // Handle potential error from the repository's GetTaskByID method
	//   return err
	// }
	// if task == nil {
	//   // If GetTaskByID returns nil task (not found) and no error
	//   return errors.New("task not found for deletion")
	// }

	// Call the DeleteTask method of the repository.
	// This line returns a single value (an error or nil), matching the function's signature.
	return s.repo.DeleteTask(id)
}
