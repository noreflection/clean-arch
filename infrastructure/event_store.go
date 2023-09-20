package infrastructure

import (
	"database/sql"
)

// EventStore represents the event store.
type EventStore struct {
	db *sql.DB
}

// NewEventStore creates a new event store.
func NewEventStore(db *sql.DB) *EventStore {
	return &EventStore{db}
}

// Implement event store functions for order and customer events here
