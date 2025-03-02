package events

import (
	"time"
)

// OrderCreatedEvent represents an event when an order is created
type OrderCreatedEvent struct {
	ID        string
	Product   string
	Quantity  int
	CreatedAt time.Time
}

// NewOrderCreatedEvent creates a new OrderCreatedEvent
func NewOrderCreatedEvent(id, product string, quantity int) *OrderCreatedEvent {
	return &OrderCreatedEvent{
		ID:        id,
		Product:   product,
		Quantity:  quantity,
		CreatedAt: time.Now(),
	}
}

// OrderUpdatedEvent represents an event when an order is updated
type OrderUpdatedEvent struct {
	ID         string
	Product    string
	Quantity   int
	CustomerID *string
	UpdatedAt  time.Time
}

// NewOrderUpdatedEvent creates a new OrderUpdatedEvent
func NewOrderUpdatedEvent(id, product string, quantity int, customerID *string) *OrderUpdatedEvent {
	return &OrderUpdatedEvent{
		ID:         id,
		Product:    product,
		Quantity:   quantity,
		CustomerID: customerID,
		UpdatedAt:  time.Now(),
	}
}

// OrderDeletedEvent represents an event when an order is deleted
type OrderDeletedEvent struct {
	ID        string
	DeletedAt time.Time
}

// NewOrderDeletedEvent creates a new OrderDeletedEvent
func NewOrderDeletedEvent(id string) *OrderDeletedEvent {
	return &OrderDeletedEvent{
		ID:        id,
		DeletedAt: time.Now(),
	}
}

// CustomerAssignedToOrderEvent represents an event when a customer is assigned to an order
type CustomerAssignedToOrderEvent struct {
	OrderID    string
	CustomerID string
	AssignedAt time.Time
}

// NewCustomerAssignedToOrderEvent creates a new CustomerAssignedToOrderEvent
func NewCustomerAssignedToOrderEvent(orderID, customerID string) *CustomerAssignedToOrderEvent {
	return &CustomerAssignedToOrderEvent{
		OrderID:    orderID,
		CustomerID: customerID,
		AssignedAt: time.Now(),
	}
}
