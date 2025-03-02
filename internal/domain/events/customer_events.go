package events

import (
	"time"
)

// CustomerCreatedEvent represents an event when a customer is created
type CustomerCreatedEvent struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}

// NewCustomerCreatedEvent creates a new CustomerCreatedEvent
func NewCustomerCreatedEvent(id, name, email string) *CustomerCreatedEvent {
	return &CustomerCreatedEvent{
		ID:        id,
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}
}

// CustomerUpdatedEvent represents an event when a customer is updated
type CustomerUpdatedEvent struct {
	ID        string
	Name      string
	Email     string
	UpdatedAt time.Time
}

// NewCustomerUpdatedEvent creates a new CustomerUpdatedEvent
func NewCustomerUpdatedEvent(id, name, email string) *CustomerUpdatedEvent {
	return &CustomerUpdatedEvent{
		ID:        id,
		Name:      name,
		Email:     email,
		UpdatedAt: time.Now(),
	}
}

// CustomerDeletedEvent represents an event when a customer is deleted
type CustomerDeletedEvent struct {
	ID        string
	DeletedAt time.Time
}

// NewCustomerDeletedEvent creates a new CustomerDeletedEvent
func NewCustomerDeletedEvent(id string) *CustomerDeletedEvent {
	return &CustomerDeletedEvent{
		ID:        id,
		DeletedAt: time.Now(),
	}
}
