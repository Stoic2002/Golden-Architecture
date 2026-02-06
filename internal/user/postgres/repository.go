package postgres

import (
	"context"
	"errors"

	"github.com/arulkarim/golden-architecture/internal/domain"
	"github.com/arulkarim/golden-architecture/internal/domain/contract"
	"github.com/arulkarim/golden-architecture/internal/domain/entity"
	"gorm.io/gorm"
)

// userRepository implements contract.UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *gorm.DB) contract.UserRepository {
	return &userRepository{db: db}
}

// Create creates a new user
func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	result := r.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		// Check for duplicate email
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return domain.ErrDuplicateEntry
		}
		return domain.ErrDatabaseOperation
	}
	return nil
}

// FindByEmail finds a user by email
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	result := r.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, domain.ErrDatabaseOperation
	}
	return &user, nil
}

// FindByID finds a user by ID
func (r *userRepository) FindByID(ctx context.Context, id uint) (*entity.User, error) {
	var user entity.User
	result := r.db.WithContext(ctx).First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, domain.ErrDatabaseOperation
	}
	return &user, nil
}
