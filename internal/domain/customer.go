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
}

func NewCustomer(name string, email string) *Customer {
	// Generate a unique ID for the customer (you can use any method you prefer).
	var customerID string

	return &Customer{
		ID:    customerID,
		Name:  name,
		Email: email,
		// Initialize other fields here if needed.
	}
}
