package event_store

import (
	"database/sql"
)

// OrderEventStore represents the event store for orders.
type OrderEventStore struct {
	db *sql.DB
}

// NewOrderEventStore creates a new event store for orders.
func NewOrderEventStore(db *sql.DB) *OrderEventStore {
	return &OrderEventStore{db}
}

// Implement order event store functions here
