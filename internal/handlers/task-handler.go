package handlers

import (
	"encoding/json"
	"net/http"
	"strings" // For path parsing

	//"github.com/AliGeeky/go-todo/internal/models"
	"github.com/AliGeeky/go-todo/internal/services"
)

// TaskHandler handles HTTP requests related to tasks.
type TaskHandler struct {
	service *services.TaskService
}

// NewTaskHandler creates a new instance of TaskHandler.
// It takes a TaskService as a dependency.
func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// GetAllTasksHandler handles GET /tasks requests.
func (h *TaskHandler) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// GetTaskByIDHandler handles GET /tasks/{id} requests.
func (h *TaskHandler) GetTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if id == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if task == nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// CreateTaskHandler handles POST /tasks requests.
func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.service.CreateTask(reqBody.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // Bad request for validation errors
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(task)
}

// UpdateTaskHandler handles PUT /tasks/{id} requests.
func (h *TaskHandler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if id == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	var reqBody struct {
		Title       string `json:"title"`
		IsCompleted bool   `json:"is_completed"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.service.UpdateTask(id, reqBody.Title, reqBody.IsCompleted)
	if err != nil {
		// Check for specific errors from service layer
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// DeleteTaskHandler handles DELETE /tasks/{id} requests.
func (h *TaskHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if id == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteTask(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content for successful deletion
}
