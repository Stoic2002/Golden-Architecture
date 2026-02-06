package handler

import "time"

// CreateTodoRequest represents the request body for creating a todo
type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required,min=1,max=255"`
	Description string `json:"description" binding:"max=1000"`
}

// UpdateTodoRequest represents the request body for updating a todo
type UpdateTodoRequest struct {
	Title       *string `json:"title" binding:"omitempty,min=1,max=255"`
	Description *string `json:"description" binding:"omitempty,max=1000"`
	Completed   *bool   `json:"completed"`
}

// TodoResponse represents the response body for a todo
type TodoResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// TodoListResponse represents the response body for a list of todos
type TodoListResponse struct {
	Todos []TodoResponse `json:"todos"`
	Total int            `json:"total"`
}

// FormatTime formats time to RFC3339
func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}
