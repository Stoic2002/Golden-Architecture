package domain

import "errors"

var (
	// ErrNotFound is returned when a resource is not found
	ErrNotFound = errors.New("resource not found")

	// ErrInvalidInput is returned when input validation fails
	ErrInvalidInput = errors.New("invalid input")

	// ErrDatabaseOperation is returned when database operation fails
	ErrDatabaseOperation = errors.New("database operation failed")

	// ErrDuplicateEntry is returned when trying to create a duplicate entry
	ErrDuplicateEntry = errors.New("duplicate entry")
)
