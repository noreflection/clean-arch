package order

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"go-cqrs/internal/domain"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, order domain.Order) (int, error) {
	var orderID int
	err := r.db.QueryRow("INSERT INTO orders (product, quantity) VALUES ($1, $2) RETURNING id", order.Product, order.Quantity).Scan(&orderID)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to create order")
	}
	return orderID, nil
}

func (r *Repository) Get(ctx context.Context, orderID int) (domain.Order, error) {
	var order domain.Order
	err := r.db.QueryRow("SELECT * FROM Orders WHERE id = $1", orderID).Scan(&order.ID, &order.CustomerId, &order.Product, &order.Quantity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Order{}, errors.Wrap(err, "failed to find order")
		}
		return domain.Order{}, errors.Wrap(err, "failed to get order")
	}
	return order, nil
}

func (r *Repository) Update(ctx context.Context, order domain.Order) error {
	query := "UPDATE Orders SET product = $1, quantity = $2 WHERE id = $3"
	_, err := r.db.Exec(query, order.Product, order.Quantity, order.ID)
	if err != nil {
		return errors.Wrap(err, "failed to update order")
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare("DELETE FROM Orders WHERE id = $1")
	if err != nil {
		return errors.Wrap(err, "failed to prepare statement for deleting order")
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.Wrap(err, "failed to delete order")
	}
	return nil
}
