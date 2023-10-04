package command_database

import (
	"database/sql"
	"errors"
	"log"
)

var (
	ErrCustomerNotFound = errors.New("customer not found")
)

type CustomerCommandDB struct {
	db *sql.DB
}

// NewCustomerCommandDB creates a new instance of CustomerCommandDB
func NewCustomerCommandDB(db *sql.DB) *CustomerCommandDB {
	return &CustomerCommandDB{
		db: db,
	}
}

// Create inserts a new customer into the database
func (c *CustomerCommandDB) Create(name, email string) (int64, error) {
	query := `INSERT INTO customers (name, email) VALUES (?, ?)`
	result, err := c.db.Exec(query, name, email)
	if err != nil {
		log.Println("Error inserting customer:", err)
		return 0, err
	}

	return result.LastInsertId()
}

// Update updates an existing customer's details
func (c *CustomerCommandDB) Update(id int64, name, email string) error {
	query := `UPDATE customers SET name = ?, email = ? WHERE id = ?`
	_, err := c.db.Exec(query, name, email, id)
	if err != nil {
		log.Println("Error updating customer:", err)
		return err
	}

	return nil
}

// Delete removes a customer from the database
func (c *CustomerCommandDB) Delete(id int64) error {
	query := `DELETE FROM customers WHERE id = ?`
	_, err := c.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting customer:", err)
		return err
	}

	return nil
}
