package handler

import (
	"errors"

	"github.com/arulkarim/golden-architecture/internal/infrastructure/auth"
	"github.com/arulkarim/golden-architecture/internal/user"
	"github.com/arulkarim/golden-architecture/pkg/response"
	"github.com/gin-gonic/gin"
)

// Handler handles HTTP requests for user/auth
type Handler struct {
	service *user.Service
}

// NewHandler creates a new user handler
func NewHandler(service *user.Service) *Handler {
	return &Handler{service: service}
}

// Register handles POST /api/v1/auth/register
func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", err.Error())
		return
	}

	input := user.RegisterInput{
		Email:    req.Email,
		Password: req.Password,
	}

	result, err := h.service.Register(c.Request.Context(), input)
	if err != nil {
		if errors.Is(err, user.ErrEmailAlreadyExists) {
			response.BadRequest(c, "Registration failed", "Email already exists")
			return
		}
		response.InternalServerError(c, "Registration failed", err.Error())
		return
	}

	resp := AuthResponse{
		Token: result.Token,
		User: UserResponse{
			ID:        result.User.ID,
			Email:     result.User.Email,
			CreatedAt: FormatTime(result.User.CreatedAt),
			UpdatedAt: FormatTime(result.User.UpdatedAt),
		},
	}

	response.Created(c, "User registered successfully", resp)
}

// Login handles POST /api/v1/auth/login
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", err.Error())
		return
	}

	input := user.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	}

	result, err := h.service.Login(c.Request.Context(), input)
	if err != nil {
		if errors.Is(err, user.ErrInvalidCredentials) {
			response.Error(c, 401, "Login failed", "Invalid email or password")
			return
		}
		response.InternalServerError(c, "Login failed", err.Error())
		return
	}

	resp := AuthResponse{
		Token: result.Token,
		User: UserResponse{
			ID:        result.User.ID,
			Email:     result.User.Email,
			CreatedAt: FormatTime(result.User.CreatedAt),
			UpdatedAt: FormatTime(result.User.UpdatedAt),
		},
	}

	response.OK(c, "Login successful", resp)
}

// Profile handles GET /api/v1/auth/profile
func (h *Handler) Profile(c *gin.Context) {
	userID, ok := auth.GetUserIDFromContext(c)
	if !ok {
		response.Error(c, 401, "Unauthorized", "User not found in context")
		return
	}

	u, err := h.service.GetProfile(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			response.NotFound(c, "User not found")
			return
		}
		response.InternalServerError(c, "Failed to get profile", err.Error())
		return
	}

	resp := UserResponse{
		ID:        u.ID,
		Email:     u.Email,
		CreatedAt: FormatTime(u.CreatedAt),
		UpdatedAt: FormatTime(u.UpdatedAt),
	}

	response.OK(c, "Profile retrieved successfully", resp)
}
