package event_store

import (
	"context"
	"go-cqrs/internal/domain/events"
	"sync"
)

// EventStore is an interface for storing and retrieving events.
type EventStore interface {
	StoreEvent(ctx context.Context, event events.Event) error
	GetEvents(ctx context.Context, eventType string) ([]events.Event, error)
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

// GetEvents returns events filtered by type.
func (s *InMemoryEventStore) GetEvents(ctx context.Context, eventType string) ([]events.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var filteredEvents []events.Event
	for _, event := range s.events {
		if eventType == "" || event.EventType() == eventType {
			filteredEvents = append(filteredEvents, event)
		}
	}

	return filteredEvents, nil
}

// NewInMemoryEventStore creates a new in-memory event store.
func NewInMemoryEventStore(eventType string) EventStore {
	return &InMemoryEventStore{
		storeType: eventType,
		events:    make([]events.Event, 0),
	}
}

// NewEventStore creates a new event store (currently uses in-memory implementation).
// This is maintained for backward compatibility.
func NewEventStore(eventType string) EventStore {
	return NewInMemoryEventStore(eventType)
}
