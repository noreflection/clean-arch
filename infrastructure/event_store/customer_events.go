package event_store

import (
	"database/sql"
)

// CustomerEventStore represents the event store for customers.
type CustomerEventStore struct {
	db *sql.DB
}

// NewCustomerEventStore creates a new event store for customers.
func NewCustomerEventStore(db *sql.DB) *CustomerEventStore {
	return &CustomerEventStore{db}
}

// Implement customer event store functions here
