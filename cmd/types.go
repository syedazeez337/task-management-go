package cmd

import (
	"sync"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskManager struct {
	tasks  map[int]Task
	mutex  sync.RWMutex
	lastID int
}

type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
