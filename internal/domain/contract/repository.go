package contract

import (
	"context"

	"github.com/arulkarim/golden-architecture/internal/domain/entity"
)

// TodoRepository defines the interface for todo data operations
type TodoRepository interface {
	// Create creates a new todo item
	Create(ctx context.Context, todo *entity.Todo) error

	// FindByID finds a todo by its ID
	FindByID(ctx context.Context, id uint) (*entity.Todo, error)

	// FindAll retrieves all todos
	FindAll(ctx context.Context) ([]entity.Todo, error)

	// Update updates an existing todo
	Update(ctx context.Context, todo *entity.Todo) error

	// Delete deletes a todo by its ID
	Delete(ctx context.Context, id uint) error
}

// UserRepository defines the interface for user data operations
type UserRepository interface {
	// Create creates a new user
	Create(ctx context.Context, user *entity.User) error

	// FindByEmail finds a user by email
	FindByEmail(ctx context.Context, email string) (*entity.User, error)

	// FindByID finds a user by ID
	FindByID(ctx context.Context, id uint) (*entity.User, error)
}
