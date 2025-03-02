package domain

import (
	"regexp"

	domainerrors "go-cqrs/internal/domain/errors"
)

// Customer represents a customer in the domain
type Customer struct {
	ID    int
	Name  string
	Email string
	//CreatedAt time.Time
	//UpdatedAt time.Time
}

func NewCustomer(name string, email string) (*Customer, error) {
	customer := &Customer{
		Name:  name,
		Email: email,
	}

	if err := customer.Validate(); err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *Customer) Validate() error {
	if c.Name == "" {
		return domainerrors.NewValidationError("customer name cannot be empty")
	}

	if !isValidEmail(c.Email) {
		return domainerrors.NewValidationError("invalid email format")
	}

	return nil
}

func (c *Customer) Update(name string, email string) error {
	c.Name = name
	c.Email = email
	return c.Validate()
}

// isValidEmail validates email format using simple regex
func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}
