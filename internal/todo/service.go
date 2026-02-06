package todo

import (
	"context"
	"errors"

	"github.com/arulkarim/golden-architecture/internal/domain"
	"github.com/arulkarim/golden-architecture/internal/domain/contract"
	"github.com/arulkarim/golden-architecture/internal/domain/entity"
)

// Service provides todo business logic
type Service struct {
	repo contract.TodoRepository
}

// NewService creates a new todo service
func NewService(repo contract.TodoRepository) *Service {
	return &Service{repo: repo}
}

// CreateTodoInput represents input for creating a todo
type CreateTodoInput struct {
	Title       string
	Description string
}

// UpdateTodoInput represents input for updating a todo
type UpdateTodoInput struct {
	Title       *string
	Description *string
	Completed   *bool
}

// Create creates a new todo
func (s *Service) Create(ctx context.Context, input CreateTodoInput) (*entity.Todo, error) {
	if input.Title == "" {
		return nil, domain.ErrInvalidInput
	}

	todo := &entity.Todo{
		Title:       input.Title,
		Description: input.Description,
		Completed:   false,
	}

	if err := s.repo.Create(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

// GetByID retrieves a todo by ID
func (s *Service) GetByID(ctx context.Context, id uint) (*entity.Todo, error) {
	if id == 0 {
		return nil, domain.ErrInvalidInput
	}

	todo, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// GetAll retrieves all todos
func (s *Service) GetAll(ctx context.Context) ([]entity.Todo, error) {
	return s.repo.FindAll(ctx)
}

// Update updates an existing todo
func (s *Service) Update(ctx context.Context, id uint, input UpdateTodoInput) (*entity.Todo, error) {
	if id == 0 {
		return nil, domain.ErrInvalidInput
	}

	todo, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if input.Title != nil {
		todo.Title = *input.Title
	}
	if input.Description != nil {
		todo.Description = *input.Description
	}
	if input.Completed != nil {
		todo.Completed = *input.Completed
	}

	if err := s.repo.Update(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

// Delete deletes a todo by ID
func (s *Service) Delete(ctx context.Context, id uint) error {
	if id == 0 {
		return domain.ErrInvalidInput
	}

	return s.repo.Delete(ctx, id)
}

// IsNotFound checks if error is a not found error
func IsNotFound(err error) bool {
	return errors.Is(err, domain.ErrNotFound)
}

// IsInvalidInput checks if error is an invalid input error
func IsInvalidInput(err error) bool {
	return errors.Is(err, domain.ErrInvalidInput)
}
