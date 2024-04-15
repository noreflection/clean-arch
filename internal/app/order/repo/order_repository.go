package repo

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"go-cqrs/internal/domain"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository { //todo: check why cycle dep happens here
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order domain.Order) (int, error) {
	var orderID int
	err := r.db.QueryRow("INSERT INTO orders (product, quantity) VALUES ($1, $2) RETURNING id", order.Product, order.Quantity).Scan(&orderID)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to create order")
	}
	fmt.Printf("id:%d", orderID)
	return orderID, nil

}

//func (r *OrderRepository) Create(order domain.Order) (string, error) {
//	// Insert a new order into the database and return the ID of the newly created order.
//	stmt, err := r.db.Prepare("INSERT INTO orders (...) VALUES (...)")
//	if err != nil {
//		return "", errors.Wrap(err, "failed to prepare statement for creating order")
//	}
//	defer stmt.Close()
//
//	_, err = stmt.Exec( /* values for other columns */ )
//	if err != nil {
//		return "", errors.Wrap(err, "failed to create order")
//	}
//	// Retrieve the auto-generated ID of the newly created order
//	orderID, err := r.getLastInsertedOrderID()
//	if err != nil {
//		return "", errors.Wrap(err, "failed to retrieve order ID")
//	}
//	return orderID, nil
//}

func (r *OrderRepository) Get(orderID int) (domain.Order, error) {
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

func (r *OrderRepository) Update(order domain.Order) error {
	if err := r.checkOrderExists(order.ID); err != nil {
		return err
	}

	query := "UPDATE Orders SET product = $1, quantity = $2 WHERE id = $3"
	_, err := r.db.Exec(query, order.Product, order.Quantity, order.ID)
	if err != nil {
		return errors.Wrap(err, "failed to update order")
	}
	return nil
}

func (r *OrderRepository) Delete(orderID int) error {
	if err := r.checkOrderExists(orderID); err != nil {
		return err
	}

	stmt, err := r.db.Prepare("DELETE FROM Orders WHERE id = $1")
	if err != nil {
		return errors.Wrap(err, "failed to prepare statement for deleting order")
	}
	defer stmt.Close()

	_, err = stmt.Exec(orderID)
	if err != nil {
		return errors.Wrap(err, "failed to delete order")
	}
	return nil
}

func (r *OrderRepository) checkOrderExists(orderID int) error {
	_, err := r.Get(orderID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("order with ID %d not found", orderID)
		}
		return errors.Wrap(err, "failed to check if order exists")
	}
	return nil
}
