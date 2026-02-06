package postgres

import (
	"context"
	"errors"

	"github.com/arulkarim/golden-architecture/internal/domain"
	"github.com/arulkarim/golden-architecture/internal/domain/contract"
	"github.com/arulkarim/golden-architecture/internal/domain/entity"
	"gorm.io/gorm"
)

// todoRepository implements contract.TodoRepository
type todoRepository struct {
	db *gorm.DB
}

// NewTodoRepository creates a new TodoRepository instance
func NewTodoRepository(db *gorm.DB) contract.TodoRepository {
	return &todoRepository{db: db}
}

// Create creates a new todo item
func (r *todoRepository) Create(ctx context.Context, todo *entity.Todo) error {
	result := r.db.WithContext(ctx).Create(todo)
	if result.Error != nil {
		return domain.ErrDatabaseOperation
	}
	return nil
}

// FindByID finds a todo by its ID
func (r *todoRepository) FindByID(ctx context.Context, id uint) (*entity.Todo, error) {
	var todo entity.Todo
	result := r.db.WithContext(ctx).First(&todo, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, domain.ErrDatabaseOperation
	}
	return &todo, nil
}

// FindAll retrieves all todos
func (r *todoRepository) FindAll(ctx context.Context) ([]entity.Todo, error) {
	var todos []entity.Todo
	result := r.db.WithContext(ctx).Order("created_at DESC").Find(&todos)
	if result.Error != nil {
		return nil, domain.ErrDatabaseOperation
	}
	return todos, nil
}

// Update updates an existing todo
func (r *todoRepository) Update(ctx context.Context, todo *entity.Todo) error {
	result := r.db.WithContext(ctx).Save(todo)
	if result.Error != nil {
		return domain.ErrDatabaseOperation
	}
	return nil
}

// Delete deletes a todo by its ID
func (r *todoRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&entity.Todo{}, id)
	if result.Error != nil {
		return domain.ErrDatabaseOperation
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}
