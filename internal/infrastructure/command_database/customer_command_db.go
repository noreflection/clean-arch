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

func NewCustomerCommandDB(db *sql.DB) *CustomerCommandDB {
	return &CustomerCommandDB{
		db: db,
	}
}

func (c *CustomerCommandDB) Create(name, email string) (int64, error) {
	query := `INSERT INTO customers (name, email) VALUES (?, ?)`
	result, err := c.db.Exec(query, name, email)
	if err != nil {
		log.Println("Error inserting customer:", err)
		return 0, err
	}
	return result.LastInsertId()
}

func (c *CustomerCommandDB) Update(id int64, name, email string) error {
	query := `UPDATE customers SET name = ?, email = ? WHERE id = ?`
	_, err := c.db.Exec(query, name, email, id)
	if err != nil {
		log.Println("Error updating customer:", err)
		return err
	}
	return nil
}

func (c *CustomerCommandDB) Delete(id int64) error {
	query := `DELETE FROM customers WHERE id = ?`
	_, err := c.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting customer:", err)
		return err
	}
	return nil
}
