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

func (e *DomainError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *DomainError) Unwrap() error {
	return e.Err
}

func NewNotFoundError(entity string, id interface{}) *DomainError {
	return &DomainError{
		Code:    ErrorCodeNotFound,
		Message: fmt.Sprintf("%s with ID %v not found", entity, id),
	}
}

func NewValidationError(message string) *DomainError {
	return &DomainError{
		Code:    ErrorCodeValidation,
		Message: message,
	}
}

func NewInvalidInputError(message string) *DomainError {
	return &DomainError{
		Code:    ErrorCodeInvalidInput,
		Message: message,
	}
}

func NewDatabaseError(err error, operation string) *DomainError {
	return &DomainError{
		Code:    ErrorCodeDatabaseError,
		Message: fmt.Sprintf("database error during %s", operation),
		Err:     err,
	}
}
