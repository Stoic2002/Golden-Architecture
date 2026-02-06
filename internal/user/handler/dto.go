package handler

import "time"

// RegisterRequest represents the request body for user registration
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginRequest represents the request body for user login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UserResponse represents the response body for a user
type UserResponse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// AuthResponse represents the response body for authentication
type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

// FormatTime formats time to RFC3339
func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}
