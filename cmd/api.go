package cmd

import (
	"encoding/json"
	"net/http"
	"time"
)

// instance of taskmanager
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make(map[int]Task),
	}
}

// error func
func (e APIError) Error() string {
	return e.Message
}

func (tm *TaskManager) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	tm.lastID++
	task.ID = tm.lastID
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	task.Status = "pending"

	tm.tasks[task.ID] = task

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// GetTasks returns all tasks
func (tm *TaskManager) GetTasks(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    tm.mutex.RLock()
    defer tm.mutex.RUnlock()

    tasks := make([]Task, 0, len(tm.tasks))
    for _, task := range tm.tasks {
        tasks = append(tasks, task)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}
