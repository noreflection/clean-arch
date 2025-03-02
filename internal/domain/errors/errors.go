package errors

import (
	"fmt"
)

// ErrorCode represents a machine-readable error code
type ErrorCode string

const (
	ErrorCodeNotFound      ErrorCode = "NOT_FOUND"
	ErrorCodeValidation    ErrorCode = "VALIDATION_ERROR"
	ErrorCodeUnauthorized  ErrorCode = "UNAUTHORIZED"
	ErrorCodeDatabaseError ErrorCode = "DATABASE_ERROR"
	ErrorCodeInvalidInput  ErrorCode = "INVALID_INPUT"
)

// DomainError represents an error in the domain layer
type DomainError struct {
	Code    ErrorCode
	Message string
	Err     error
}

// Error returns the error message
func (e *DomainError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap returns the wrapped error
func (e *DomainError) Unwrap() error {
	return e.Err
}

// NewNotFoundError creates a new not found error
func NewNotFoundError(entity string, id interface{}) *DomainError {
	return &DomainError{
		Code:    ErrorCodeNotFound,
		Message: fmt.Sprintf("%s with ID %v not found", entity, id),
	}
}

// NewValidationError creates a new validation error
func NewValidationError(message string) *DomainError {
	return &DomainError{
		Code:    ErrorCodeValidation,
		Message: message,
	}
}

// NewInvalidInputError creates a new invalid input error
func NewInvalidInputError(message string) *DomainError {
	return &DomainError{
		Code:    ErrorCodeInvalidInput,
		Message: message,
	}
}

// NewDatabaseError wraps a database error
func NewDatabaseError(err error, operation string) *DomainError {
	return &DomainError{
		Code:    ErrorCodeDatabaseError,
		Message: fmt.Sprintf("database error during %s", operation),
		Err:     err,
	}
}
