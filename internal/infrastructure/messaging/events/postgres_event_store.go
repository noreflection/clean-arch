package event_store

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go-cqrs/internal/domain/events"
	"go-cqrs/internal/infrastructure/logger"
)

// PostgresEventStore is a PostgreSQL implementation of the EventStore interface
type PostgresEventStore struct {
	db     *sql.DB
	logger logger.Logger
	name   string
}

// NewPostgresEventStore creates a new PostgreSQL-based event store
func NewPostgresEventStore(db *sql.DB, name string, logger logger.Logger) EventStore {
	return &PostgresEventStore{
		db:     db,
		logger: logger,
		name:   name,
	}
}

// StoreEvent stores an event in the PostgreSQL event store
func (s *PostgresEventStore) StoreEvent(ctx context.Context, event events.Event) error {
	// Serialize event to JSON
	eventData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	// Insert event into database
	_, err = s.db.ExecContext(ctx,
		`INSERT INTO events (event_type, occurred_at, event_data) VALUES ($1, $2, $3)`,
		event.EventType(), event.OccurredAt(), eventData)
	if err != nil {
		return fmt.Errorf("failed to store event: %w", err)
	}

	s.logger.Debug("Event stored successfully",
		logger.String("event_type", event.EventType()),
		logger.String("event_store", s.name))

	return nil
}

// GetEvents retrieves events by type from the event store
func (s *PostgresEventStore) GetEvents(ctx context.Context, eventType string) ([]events.Event, error) {
	// Query events from database
	rows, err := s.db.QueryContext(ctx,
		`SELECT event_type, event_data FROM events WHERE event_type = $1 ORDER BY occurred_at ASC`,
		eventType)
	if err != nil {
		return nil, fmt.Errorf("failed to query events: %w", err)
	}
	defer rows.Close()

	// Process query results
	var events []events.Event
	for rows.Next() {
		var eventType string
		var eventData []byte

		if err := rows.Scan(&eventType, &eventData); err != nil {
			return nil, fmt.Errorf("failed to scan event row: %w", err)
		}

		// Deserialize event based on type
		event, err := deserializeEvent(eventType, eventData)
		if err != nil {
			s.logger.Error("Failed to deserialize event",
				logger.String("event_type", eventType),
				logger.Error(err))
			continue
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating event rows: %w", err)
	}

	return events, nil
}

// deserializeEvent deserializes an event based on its type
func deserializeEvent(eventType string, data []byte) (events.Event, error) {
	switch eventType {
	case events.CustomerCreatedEventType:
		var event events.CustomerCreatedEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &event, nil
	case events.CustomerUpdatedEventType:
		var event events.CustomerUpdatedEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &event, nil
	case events.CustomerDeletedEventType:
		var event events.CustomerDeletedEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &event, nil
	case events.OrderCreatedEventType:
		var event events.OrderCreatedEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &event, nil
	case events.OrderUpdatedEventType:
		var event events.OrderUpdatedEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &event, nil
	case events.OrderDeletedEventType:
		var event events.OrderDeletedEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &event, nil
	case events.CustomerAssignedToOrderEventType:
		var event events.CustomerAssignedToOrderEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &event, nil
	default:
		return nil, fmt.Errorf("unknown event type: %s", eventType)
	}
}
