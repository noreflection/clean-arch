package event_store

import (
	"context"
	"sync"
)

// EventStore is an interface for storing events.
type EventStore interface {
	StoreEvent(ctx context.Context, event interface{}) error
}

// OrderCreatedEvent represents an event where an order is created.
type OrderCreatedEvent struct {
	ID       string
	Product  string
	Quantity int
}

// NewOrderCreatedEvent creates a new order created event.
func NewOrderCreatedEvent(id, product string, quantity int) OrderCreatedEvent {
	return OrderCreatedEvent{id, product, quantity}
}

//// CustomerCreatedEvent represents an event where a customer is created.
//type CustomerCreatedEvent struct {
//	ID    string
//	Name  string
//	Email string
//}
//
//// NewCustomerCreatedEvent creates a new customer created event.
//func NewCustomerCreatedEvent(id, name, email string) CustomerCreatedEvent {
//	return CustomerCreatedEvent{id, name, email}
//}

// InMemoryEventStore is an in-memory implementation of the EventStore interface.
type InMemoryEventStore struct {
	events []interface{}
	mu     sync.RWMutex
}

func (s *InMemoryEventStore) StoreEvent(ctx context.Context, event interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, event)
	return nil
}

// InMemoryCustomerEventStore is an in-memory implementation of the CustomerEventStore interface.
type InMemoryCustomerEventStore struct {
	events []interface{}
	mu     sync.RWMutex
}

func (s *InMemoryCustomerEventStore) StoreEvent(ctx context.Context, event interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, event)
	return nil
}

// NewEventStore creates a new event store based on the specified type.
func NewEventStore(eventType string) EventStore {
	switch eventType {
	case "order":
		return &InMemoryEventStore{events: make([]interface{}, 0)}
	case "customer":
		return &InMemoryCustomerEventStore{events: make([]interface{}, 0)}
	default:
		panic("Unsupported event type")
	}
}
