package event_store

import (
	"context"
	"database/sql"
	"time"
)

// CustomerCreatedEvent represents an event where a customer is created.
type CustomerCreatedEvent struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}

// CustomerEventStore represents the event store for customers.
type CustomerEventStore struct {
	db *sql.DB
}

// NewCustomerEventStore creates a new event store for customers.
func NewCustomerEventStore(db *sql.DB) *CustomerEventStore {
	return &CustomerEventStore{db}
}

// NewCustomerCreatedEvent creates a new CustomerCreatedEvent.
func NewCustomerCreatedEvent(id, name, email string) *CustomerCreatedEvent {
	return &CustomerCreatedEvent{
		ID:    id,
		Name:  name,
		Email: email,
		//Timestamp: time.Now(),
	}
}

// StoreEvent stores a customer created event in the event store.
func (s *CustomerEventStore) StoreEvent(ctx context.Context, event CustomerCreatedEvent) error {
	//query := `
	//	INSERT INTO customer_events (id, name, email, created_at)
	//	VALUES (?, ?, ?, ?)
	//`
	//_, err := s.db.ExecContext(ctx, query, event.ID, event.Name, event.Email, event.CreatedAt)
	//if err != nil {
	//	return err
	//}
	return nil
}

// RetrieveEvents retrieves all customer created events from the event store.
func (s *CustomerEventStore) RetrieveEvents(ctx context.Context) ([]CustomerCreatedEvent, error) {
	query := `
		SELECT id, name, email, created_at 
		FROM customer_events
	`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []CustomerCreatedEvent
	for rows.Next() {
		var event CustomerCreatedEvent
		if err := rows.Scan(&event.ID, &event.Name, &event.Email, &event.CreatedAt); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}
