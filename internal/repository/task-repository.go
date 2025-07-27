package repository

import (
	"sync" // For thread-safe map access

	"github.com/AliGeeky/go-todo/internal/models"
)

// TaskRepository defines the interface for task data operations.
type TaskRepository interface {
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id string) (*models.Task, error)
	CreateTask(task models.Task) error
	UpdateTask(task models.Task) error
	DeleteTask(id string) error
}

// InMemoryTaskRepository implements TaskRepository using an in-memory map.
type InMemoryTaskRepository struct {
	tasks map[string]models.Task
	mu    sync.RWMutex // Mutex for concurrent map access
}

// NewInMemoryTaskRepository creates a new instance of InMemoryTaskRepository.
func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]models.Task),
	}
}

// GetAllTasks retrieves all tasks from the in-memory store.
func (r *InMemoryTaskRepository) GetAllTasks() ([]models.Task, error) {
	r.mu.RLock()         // Read lock
	defer r.mu.RUnlock() // Unlock when function returns
	var tasks []models.Task
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// GetTaskByID retrieves a task by its ID.
func (r *InMemoryTaskRepository) GetTaskByID(id string) (*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	task, ok := r.tasks[id]
	if !ok {
		return nil, nil // Or an error indicating not found
	}
	return &task, nil
}

// CreateTask adds a new task to the in-memory store.
func (r *InMemoryTaskRepository) CreateTask(task models.Task) error {
	r.mu.Lock() // Write lock
	defer r.mu.Unlock()
	r.tasks[task.ID] = task
	return nil
}

// UpdateTask updates an existing task in the in-memory store.
func (r *InMemoryTaskRepository) UpdateTask(task models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.tasks[task.ID]; !ok {
		return nil // Or an error indicating not found
	}
	r.tasks[task.ID] = task
	return nil
}

// DeleteTask deletes a task from the in-memory store.
func (r *InMemoryTaskRepository) DeleteTask(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.tasks, id)
	return nil
}
