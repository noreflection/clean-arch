package customer

import (
	"cqrs-web-api/internal/domain"
	"database/sql"
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

func (s *CustomerService) CreateCustomer(customer domain.Customer) (*domain.Customer, error) {
	// Implement the CreateCustomer method with PostgreSQL database operations
	_, err := s.db.Exec("INSERT INTO customers (id, name) VALUES ($1, $2)", customer.ID, customer.Name)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (s *CustomerService) GetCustomer(customerID string) (*domain.Customer, error) {
	// Implement the GetCustomer method with PostgreSQL database operations
	var customer domain.Customer
	err := s.db.QueryRow("SELECT id, name FROM customers WHERE id = $1", customerID).Scan(&customer.ID, &customer.Name)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

// Add other customer-related service methods here
