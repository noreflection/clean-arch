package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-cqrs/internal/domain"
)

// OrderRepository implements app.OrderRepository
type OrderRepository struct {
	db *sql.DB
}

// NewOrderRepository creates a new OrderRepository
func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// Create inserts a new order into the database
func (r *OrderRepository) Create(ctx context.Context, order domain.Order) (int, error) {
	var orderID int
	
	query := "INSERT INTO orders (product, quantity) VALUES ($1, $2) RETURNING id"
	if order.CustomerID != nil {
		query = "INSERT INTO orders (customer_id, product, quantity) VALUES ($1, $2, $3) RETURNING id"
		err := r.db.QueryRowContext(ctx, query, *order.CustomerID, order.Product, order.Quantity).Scan(&orderID)
		if err != nil {
			return 0, errors.New("failed to create order: " + err.Error())
		}
	} else {
		err := r.db.QueryRowContext(ctx, query, order.Product, order.Quantity).Scan(&orderID)
		if err != nil {
			return 0, errors.New("failed to create order: " + err.Error())
		}
	}
	
	return orderID, nil
}

// GetByID retrieves an order by its ID
func (r *OrderRepository) GetByID(ctx context.Context, id int) (*domain.Order, error) {
	var order domain.Order
	var customerID sql.NullInt64
	
	err := r.db.QueryRowContext(ctx, "SELECT id, customer_id, product, quantity FROM orders WHERE id = $1", id).
		Scan(&order.ID, &customerID, &order.Product, &order.Quantity)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found, return nil without error
		}
		return nil, errors.New("failed to get order: " + err.Error())
	}
	
	if customerID.Valid {
		custID := int(customerID.Int64)
		order.CustomerID = &custID
	}
	
	return &order, nil
}

// GetByCustomerID retrieves all orders for a customer
func (r *OrderRepository) GetByCustomerID(ctx context.Context, customerID int) ([]domain.Order, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, customer_id, product, quantity FROM orders WHERE customer_id = $1", customerID)
	if err != nil {
		return nil, errors.New("failed to get orders by customer: " + err.Error())
	}
	defer rows.Close()
	
	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		var customerID sql.NullInt64
		
		err := rows.Scan(&order.ID, &customerID, &order.Product, &order.Quantity)
		if err != nil {
			return nil, errors.New("failed to scan order row: " + err.Error())
		}
		
		if customerID.Valid {
			custID := int(customerID.Int64)
			order.CustomerID = &custID
		}
		
		orders = append(orders, order)
	}
	
	if err = rows.Err(); err != nil {
		return nil, errors.New("error iterating order rows: " + err.Error())
	}
	
	return orders, nil
}

// Update updates an existing order
func (r *OrderRepository) Update(ctx context.Context, order domain.Order) error {
	var err error
	
	if order.CustomerID != nil {
		_, err = r.db.ExecContext(ctx, 
			"UPDATE orders SET customer_id = $1, product = $2, quantity = $3 WHERE id = $4",
			*order.CustomerID, order.Product, order.Quantity, order.ID)
	} else {
		_, err = r.db.ExecContext(ctx, 
			"UPDATE orders SET customer_id = NULL, product = $1, quantity = $2 WHERE id = $3",
			order.Product, order.Quantity, order.ID)
	}
	
	if err != nil {
		return errors.New("failed to update order: " + err.Error())
	}
	
	return nil
}

// Delete removes an order
func (r *OrderRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM orders WHERE id = $1", id)
	if err != nil {
		return errors.New("failed to delete order: " + err.Error())
	}
	
	return nil
}

// List retrieves orders with pagination
func (r *OrderRepository) List(ctx context.Context, limit, offset int) ([]domain.Order, error) {
	rows, err := r.db.QueryContext(ctx, 
		"SELECT id, customer_id, product, quantity FROM orders LIMIT $1 OFFSET $2",
		limit, offset)
	if err != nil {
		return nil, errors.New("failed to list orders: " + err.Error())
	}
	defer rows.Close()
	
	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		var customerID sql.NullInt64
		
		err := rows.Scan(&order.ID, &customerID, &order.Product, &order.Quantity)
		if err != nil {
			return nil, errors.New("failed to scan order row: " + err.Error())
		}
		
		if customerID.Valid {
			custID := int(customerID.Int64)
			order.CustomerID = &custID
		}
		
		orders = append(orders, order)
	}
	
	if err = rows.Err(); err != nil {
		return nil, errors.New("error iterating order rows: " + err.Error())
	}
	
	return orders, nil
}
