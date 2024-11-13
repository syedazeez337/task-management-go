package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/syedazeez337/task-management-go/cmd"
)

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next(w, r)
		log.Printf("Completed in %v", time.Since(start))
	}
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// task manager instance
	var taskManager = cmd.NewTaskManager()

	// Register routes with middleware
	mux.HandleFunc("/tasks", loggingMiddleware(taskManager.GetTasks))
	mux.HandleFunc("/tasks/create", loggingMiddleware(taskManager.CreateTask))

	// Start the server
	fmt.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
