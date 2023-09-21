package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID    string
	Name  string
	Email string
	// Add other fields specific to customers
}

// NewCustomer creates a new customer instance with the given name.
func NewCustomer(name string, email string) *Customer {
	// Generate a unique ID for the customer (you can use any method you prefer).
	customerID := generateUniqueID()

	return &Customer{
		ID:    customerID,
		Name:  name,
		Email: email,
		// Initialize other fields here if needed.
	}
}

// generateUniqueID generates a unique identifier for the customer.
func generateUniqueID() string {
	// You can use a library like "github.com/google/uuid" to generate UUIDs.
	// Example:
	uniqueID := uuid.New().String()
	return uniqueID
}
