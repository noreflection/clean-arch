package domain

import (
	"errors"
	"regexp"
)

// Customer represents a customer in the domain
type Customer struct {
	ID    int
	Name  string
	Email string
	//CreatedAt time.Time
	//UpdatedAt time.Time
}

// NewCustomer creates a new Customer entity with validation
func NewCustomer(name string, email string) (*Customer, error) {
	if name == "" {
		return nil, errors.New("customer name cannot be empty")
	}
	
	if !isValidEmail(email) {
		return nil, errors.New("invalid email format")
	}
	
	return &Customer{
		Name:  name,
		Email: email,
	}, nil
}

// isValidEmail validates email format using simple regex
func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}
