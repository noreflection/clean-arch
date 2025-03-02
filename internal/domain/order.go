package domain

import (
	domainerrors "go-cqrs/internal/domain/errors"
)

// Order represents an order in the domain
type Order struct {
	ID         int
	CustomerID *int // Using pointer instead of sql.NullInt64 to represent optional value
	Product    string
	Quantity   int
	// Could add other domain-related fields like:
	// Status    OrderStatus
	// CreatedAt time.Time
}

// OrderStatus represents the current state of an order
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusConfirmed OrderStatus = "CONFIRMED"
	OrderStatusShipped   OrderStatus = "SHIPPED"
	OrderStatusDelivered OrderStatus = "DELIVERED"
	OrderStatusCancelled OrderStatus = "CANCELLED"
)

// NewOrder creates a new Order entity with validation
func NewOrder(product string, quantity int) (*Order, error) {
	order := &Order{
		Product:  product,
		Quantity: quantity,
	}

	if err := order.Validate(); err != nil {
		return nil, err
	}

	return order, nil
}

// Validate validates the order entity
func (o *Order) Validate() error {
	if o.Product == "" {
		return domainerrors.NewValidationError("product cannot be empty")
	}

	if o.Quantity <= 0 {
		return domainerrors.NewValidationError("quantity must be greater than zero")
	}

	return nil
}

// AssignCustomer assigns a customer to this order with validation
func (o *Order) AssignCustomer(customerID int) error {
	if customerID <= 0 {
		return domainerrors.NewValidationError("customer ID must be greater than zero")
	}

	o.CustomerID = &customerID
	return nil
}

// Update updates the order fields with validation
func (o *Order) Update(product string, quantity int) error {
	o.Product = product
	o.Quantity = quantity
	return o.Validate()
}
