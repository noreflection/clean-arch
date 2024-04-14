package domain

import (
	"database/sql"
	"github.com/google/uuid"
)

type Order struct {
	ID       string // Changed ID type to int to match the provided id in NewOrder function
	Product  string
	Quantity int
}

func NewOrder(id string, product string, quantity int) (*Order, error) {
	//customerID := generateUniqueID()

	return &Order{
		ID:       id,
		Product:  product,
		Quantity: quantity,
	}, nil
}

func generateUniqueID() string {
	// You can use a library like "github.com/google/uuid" to generate UUIDs.
	uniqueID := uuid.New().String()
	return uniqueID
}

func SaveOrder(db *sql.DB, order *Order) error {
	// Prepare the SQL statement
	stmt, err := db.Prepare("INSERT INTO orders (id, product, quantity) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(order.ID, order.Product, order.Quantity)
	if err != nil {
		return err
	}

	return nil
}
