package customer

import (
	"database/sql"
	"go-cqrs/internal/domain"
)

// CustomerService represents the service for customers.
type CustomerService struct {
	db *sql.DB
}

func NewService(db *sql.DB) *CustomerService {
	return &CustomerService{
		db: db,
	}
}

// CreateCustomer creates a new customer in the database.
func (s *CustomerService) CreateCustomer(customer domain.Customer) (*domain.Customer, error) {
	// Implement the CreateCustomer method with PostgreSQL database operations
	_, err := s.db.Exec("INSERT INTO customers (id, name) VALUES ($1, $2)", customer.ID, customer.Name)
	if err != nil {
		return nil, err
	}

	// Return the created customer, assuming the database operation was successful.
	return &customer, nil
}

// GetCustomer retrieves a customer by ID from the database.
func (s *CustomerService) GetCustomer(customerID string) (*domain.Customer, error) {
	// Implement the GetCustomer method with PostgreSQL database operations
	var customer domain.Customer
	err := s.db.QueryRow("SELECT id, name FROM customers WHERE id = $1", customerID).Scan(&customer.ID, &customer.Name)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

// UpdateCustomer updates an existing customer in the database.
func (s *CustomerService) UpdateCustomer(customer domain.Customer) (*domain.Customer, error) {
	// Implement the UpdateCustomer method with PostgreSQL database operations
	_, err := s.db.Exec("UPDATE customers SET name = $2 WHERE id = $1", customer.ID, customer.Name)
	if err != nil {
		return nil, err
	}

	// Return the updated customer, assuming the database operation was successful.
	return &customer, nil
}

// DeleteCustomer deletes a customer by ID from the database.
func (s *CustomerService) DeleteCustomer(customerID string) error {
	// Implement the DeleteCustomer method with PostgreSQL database operations
	_, err := s.db.Exec("DELETE FROM customers WHERE id = $1", customerID)
	if err != nil {
		return err
	}

	return nil
}
