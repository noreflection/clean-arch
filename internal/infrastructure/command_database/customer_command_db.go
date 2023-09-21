package command_database

import (
	"database/sql"
	"go-cqrs/internal/domain"
)

// CustomerCommandDB represents the command database for customer commands.
type CustomerCommandDB struct {
	db *sql.DB
}

// NewCustomerCommandDB creates a new instance of CustomerCommandDB.
func NewCustomerCommandDB(db *sql.DB) *CustomerCommandDB {
	return &CustomerCommandDB{
		db: db,
	}
}

// CreateCustomerCommandDB handles the database operation for creating a new customer.
func (db *CustomerCommandDB) CreateCustomerCommandDB(customer domain.Customer) error {
	// Implement the database logic to create a new customer.
	_, err := db.db.Exec("INSERT INTO customers (id, name) VALUES ($1, $2)", customer.ID, customer.Name)
	if err != nil {
		return err
	}

	return nil
}

// UpdateCustomerCommandDB handles the database operation for updating an existing customer.
func (db *CustomerCommandDB) UpdateCustomerCommandDB(customer domain.Customer) error {
	// Implement the database logic to update an existing customer.
	_, err := db.db.Exec("UPDATE customers SET name = $2 WHERE id = $1", customer.ID, customer.Name)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCustomerCommandDB handles the database operation for deleting an existing customer.
func (db *CustomerCommandDB) DeleteCustomerCommandDB(customerID string) error {
	// Implement the database logic to delete an existing customer.
	_, err := db.db.Exec("DELETE FROM customers WHERE id = $1", customerID)
	if err != nil {
		return err
	}

	return nil
}
