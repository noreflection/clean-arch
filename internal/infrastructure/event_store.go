// file: event_store.go
package event_store

import (
	"context"
	"sync"
)

// EventStore is an interface for storing events related to orders.
type EventStore interface {
	StoreEvent(ctx context.Context, event OrderCreatedEvent) error
}

// CustomerEventStore is an interface for storing events related to customers.
type CustomerEventStore interface {
	StoreEvent(ctx context.Context, event CustomerCreatedEvent) error
}

// OrderCreatedEvent represents an event where an order is created.
type OrderCreatedEvent struct {
	ID       string
	Product  string
	Quantity int
}

// CustomerCreatedEvent represents an event where a customer is created.
type CustomerCreatedEvent struct {
	ID    string
	Name  string
	Email string
}

// NewOrderCreatedEvent creates a new order created event.
func NewOrderCreatedEvent(id, product string, quantity int) OrderCreatedEvent {
	return OrderCreatedEvent{id, product, quantity}
}

// NewCustomerCreatedEvent creates a new customer created event.
func NewCustomerCreatedEvent(id, name, email string) CustomerCreatedEvent {
	return CustomerCreatedEvent{id, name, email}
}

// InMemoryEventStore is an in-memory implementation of the EventStore interface.
type InMemoryEventStore struct {
	events []OrderCreatedEvent
	mu     sync.RWMutex
}

func (s *InMemoryEventStore) StoreEvent(ctx context.Context, event OrderCreatedEvent) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, event)
	return nil
}

// InMemoryCustomerEventStore is an in-memory implementation of the CustomerEventStore interface.
type InMemoryCustomerEventStore struct {
	events []CustomerCreatedEvent
	mu     sync.RWMutex
}

func (s *InMemoryCustomerEventStore) StoreEvent(ctx context.Context, event CustomerCreatedEvent) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, event)
	return nil
}
