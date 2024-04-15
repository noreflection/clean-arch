package domain

import (
	"database/sql"
)

type Order struct {
	ID         int
	CustomerId sql.NullInt64
	Product    string
	Quantity   int
}

func NewOrder(id int, product string, quantity int) (*Order, error) { // todo:
	return &Order{
		ID:       id,
		Product:  product,
		Quantity: quantity,
	}, nil
}
