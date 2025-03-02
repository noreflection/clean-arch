package event_store

import (
	"context"
	"go-cqrs/internal/domain/events"
	"sync"
)

// EventStore is an interface for storing events.
type EventStore interface {
	StoreEvent(ctx context.Context, event events.Event) error
}

// InMemoryEventStore is an in-memory implementation of the EventStore interface.
type InMemoryEventStore struct {
	storeType string
	events    []events.Event
	mu        sync.RWMutex
}

// StoreEvent stores an event in the event store.
func (s *InMemoryEventStore) StoreEvent(ctx context.Context, event events.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, event)
	return nil
}

// GetEvents returns all stored events.
func (s *InMemoryEventStore) GetEvents() []events.Event {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.events
}

// NewEventStore creates a new event store based on the specified type.
func NewEventStore(eventType string) EventStore {
	return &InMemoryEventStore{
		storeType: eventType,
		events:    make([]events.Event, 0),
	}
}
