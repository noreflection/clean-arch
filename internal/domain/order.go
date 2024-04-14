package domain

import (
	"database/sql"
)

type Order struct {
	ID         string
	CustomerId sql.NullInt64
	Product    string
	Quantity   int
}

func NewOrder(id string, product string, quantity int) (*Order, error) {
	//customerID := generateUniqueID()

	return &Order{
		ID:       id,
		Product:  product,
		Quantity: quantity,
	}, nil
}
