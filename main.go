package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/AliGeeky/go-todo/internal/handlers"
	"github.com/AliGeeky/go-todo/internal/repository"
	"github.com/AliGeeky/go-todo/internal/services"
)

func main() {
	// 1. Initialize Repository
	taskRepo := repository.NewInMemoryTaskRepository()

	// 2. Initialize Service with Repository dependency
	taskService := services.NewTaskService(taskRepo)

	// 3. Initialize Handler with Service dependency
	taskHandler := handlers.NewTaskHandler(taskService)

	// 4. Define HTTP Routes
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			taskHandler.GetAllTasksHandler(w, r)
		case http.MethodPost:
			taskHandler.CreateTaskHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		// This handler catches requests like /tasks/{id}
		// Make sure it's not just /tasks (which is handled above)
		if len(strings.Split(r.URL.Path, "/")) == 3 && strings.Split(r.URL.Path, "/")[2] != "" {
			switch r.Method {
			case http.MethodGet:
				taskHandler.GetTaskByIDHandler(w, r)
			case http.MethodPut:
				taskHandler.UpdateTaskHandler(w, r)
			case http.MethodDelete:
				taskHandler.DeleteTaskHandler(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		} else {
			http.NotFound(w, r) // If it's just /tasks/ (with trailing slash) or invalid, treat as not found
		}
	})

	// 5. Start the HTTP Server
	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil)) // Blocks until server stops or an error occurs
}
