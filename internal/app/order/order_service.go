package order

import (
	"database/sql"
)

// Service represents the order service.
type Service struct {
	db *sql.DB
}

// NewService creates a new order service.
func NewService(db *sql.DB) *Service {
	return &Service{db}
}

// Implement order-related service functions here
