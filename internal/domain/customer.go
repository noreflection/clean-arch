package domain

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID      string
	Name    string
	Surname string
	Email   string
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
