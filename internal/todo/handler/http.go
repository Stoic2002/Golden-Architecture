package handler

import (
	"strconv"

	"github.com/arulkarim/golden-architecture/internal/todo"
	"github.com/arulkarim/golden-architecture/pkg/response"
	"github.com/gin-gonic/gin"
)

// Handler handles HTTP requests for todos
type Handler struct {
	service *todo.Service
}

// NewHandler creates a new todo handler
func NewHandler(service *todo.Service) *Handler {
	return &Handler{service: service}
}

// Create handles POST /api/v1/todos
func (h *Handler) Create(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", err.Error())
		return
	}

	input := todo.CreateTodoInput{
		Title:       req.Title,
		Description: req.Description,
	}

	result, err := h.service.Create(c.Request.Context(), input)
	if err != nil {
		if todo.IsInvalidInput(err) {
			response.BadRequest(c, "Invalid input", err.Error())
			return
		}
		response.InternalServerError(c, "Failed to create todo", err.Error())
		return
	}

	resp := TodoResponse{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Completed:   result.Completed,
		CreatedAt:   FormatTime(result.CreatedAt),
		UpdatedAt:   FormatTime(result.UpdatedAt),
	}

	response.Created(c, "Todo created successfully", resp)
}

// GetAll handles GET /api/v1/todos
func (h *Handler) GetAll(c *gin.Context) {
	todos, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		response.InternalServerError(c, "Failed to get todos", err.Error())
		return
	}

	var todoResponses []TodoResponse
	for _, t := range todos {
		todoResponses = append(todoResponses, TodoResponse{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Completed:   t.Completed,
			CreatedAt:   FormatTime(t.CreatedAt),
			UpdatedAt:   FormatTime(t.UpdatedAt),
		})
	}

	resp := TodoListResponse{
		Todos: todoResponses,
		Total: len(todoResponses),
	}

	response.OK(c, "Todos retrieved successfully", resp)
}

// GetByID handles GET /api/v1/todos/:id
func (h *Handler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid todo ID", "ID must be a positive integer")
		return
	}

	result, err := h.service.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		if todo.IsNotFound(err) {
			response.NotFound(c, "Todo not found")
			return
		}
		response.InternalServerError(c, "Failed to get todo", err.Error())
		return
	}

	resp := TodoResponse{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Completed:   result.Completed,
		CreatedAt:   FormatTime(result.CreatedAt),
		UpdatedAt:   FormatTime(result.UpdatedAt),
	}

	response.OK(c, "Todo retrieved successfully", resp)
}

// Update handles PUT /api/v1/todos/:id
func (h *Handler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid todo ID", "ID must be a positive integer")
		return
	}

	var req UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", err.Error())
		return
	}

	input := todo.UpdateTodoInput{
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
	}

	result, err := h.service.Update(c.Request.Context(), uint(id), input)
	if err != nil {
		if todo.IsNotFound(err) {
			response.NotFound(c, "Todo not found")
			return
		}
		if todo.IsInvalidInput(err) {
			response.BadRequest(c, "Invalid input", err.Error())
			return
		}
		response.InternalServerError(c, "Failed to update todo", err.Error())
		return
	}

	resp := TodoResponse{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Completed:   result.Completed,
		CreatedAt:   FormatTime(result.CreatedAt),
		UpdatedAt:   FormatTime(result.UpdatedAt),
	}

	response.OK(c, "Todo updated successfully", resp)
}

// Delete handles DELETE /api/v1/todos/:id
func (h *Handler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid todo ID", "ID must be a positive integer")
		return
	}

	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		if todo.IsNotFound(err) {
			response.NotFound(c, "Todo not found")
			return
		}
		response.InternalServerError(c, "Failed to delete todo", err.Error())
		return
	}

	response.OK(c, "Todo deleted successfully", nil)
}
