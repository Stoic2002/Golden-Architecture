package user

import (
	"context"
	"errors"

	"github.com/arulkarim/golden-architecture/internal/domain"
	"github.com/arulkarim/golden-architecture/internal/domain/contract"
	"github.com/arulkarim/golden-architecture/internal/domain/entity"
	"github.com/arulkarim/golden-architecture/internal/infrastructure/auth"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserNotFound       = errors.New("user not found")
)

// Service provides user/auth business logic
type Service struct {
	repo       contract.UserRepository
	jwtManager *auth.JWTManager
}

// NewService creates a new user service
func NewService(repo contract.UserRepository, jwtManager *auth.JWTManager) *Service {
	return &Service{
		repo:       repo,
		jwtManager: jwtManager,
	}
}

// RegisterInput represents input for user registration
type RegisterInput struct {
	Email    string
	Password string
}

// LoginInput represents input for user login
type LoginInput struct {
	Email    string
	Password string
}

// AuthResult represents the result of authentication
type AuthResult struct {
	Token string
	User  *entity.User
}

// Register registers a new user
func (s *Service) Register(ctx context.Context, input RegisterInput) (*AuthResult, error) {
	// Check if email already exists
	existingUser, err := s.repo.FindByEmail(ctx, input.Email)
	if err == nil && existingUser != nil {
		return nil, ErrEmailAlreadyExists
	}
	if err != nil && !errors.Is(err, domain.ErrNotFound) {
		return nil, err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &entity.User{
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		if errors.Is(err, domain.ErrDuplicateEntry) {
			return nil, ErrEmailAlreadyExists
		}
		return nil, err
	}

	// Generate token
	token, err := s.jwtManager.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &AuthResult{
		Token: token,
		User:  user,
	}, nil
}

// Login authenticates a user
func (s *Service) Login(ctx context.Context, input LoginInput) (*AuthResult, error) {
	// Find user by email
	user, err := s.repo.FindByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	// Generate token
	token, err := s.jwtManager.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &AuthResult{
		Token: token,
		User:  user,
	}, nil
}

// GetProfile gets user profile by ID
func (s *Service) GetProfile(ctx context.Context, userID uint) (*entity.User, error) {
	user, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}
