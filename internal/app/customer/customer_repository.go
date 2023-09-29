package customer

import (
	"database/sql"
	"errors"
)

// CustomerRepository represents the repository for customers.
type CustomerRepository struct {
	db *sql.DB
}

// NewCustomerRepository creates a new customer repository.
func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db}
}

// Customer represents a customer entity in the database.
type Customer struct {
	ID   int
	Name string
	// Add other customer-related fields here
}

// CreateCustomer inserts a new customer record into the database.
func (r *CustomerRepository) CreateCustomer(customer *Customer) error {
	// Validate input, perform any necessary business logic, etc.

	// Execute SQL insert statement to create a new customer
	_, err := r.db.Exec("INSERT INTO customers (name) VALUES (?)", customer.Name)
	if err != nil {
		// Handle the error, such as rolling back a transaction or logging
		return err
	}

	// If the insert was successful, you can update the customer ID field
	// with the generated ID if needed (e.g., for auto-incrementing primary keys).

	return nil
}

// GetCustomerByID retrieves a customer record from the database by ID.
func (r *CustomerRepository) GetCustomerByID(customerID int) (*Customer, error) {
	// Execute SQL query to retrieve customer by ID
	row := r.db.QueryRow("SELECT id, name FROM customers WHERE id = ?", customerID)

	var customer Customer
	err := row.Scan(&customer.ID, &customer.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Handle the case where no customer with the given ID was found
			return nil, nil
		}
		// Handle other errors
		return nil, err
	}

	return &customer, nil
}

// Implement other customer-related repository functions here
