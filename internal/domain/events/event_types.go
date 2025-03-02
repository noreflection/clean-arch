package events

import (
	"time"
)

// Event is the base interface for all domain events
type Event interface {
	EventType() string
	OccurredAt() time.Time
}

// EventType constants
const (
	CustomerCreatedEventType         = "customer.created"
	CustomerUpdatedEventType         = "customer.updated"
	CustomerDeletedEventType         = "customer.deleted"
	OrderCreatedEventType            = "order.created"
	OrderUpdatedEventType            = "order.updated"
	OrderDeletedEventType            = "order.deleted"
	CustomerAssignedToOrderEventType = "order.customer_assigned"
)

// EventType implementations
func (e *CustomerCreatedEvent) EventType() string {
	return CustomerCreatedEventType
}

func (e *CustomerCreatedEvent) OccurredAt() time.Time {
	return e.CreatedAt
}

func (e *CustomerUpdatedEvent) EventType() string {
	return CustomerUpdatedEventType
}

func (e *CustomerUpdatedEvent) OccurredAt() time.Time {
	return e.UpdatedAt
}

func (e *CustomerDeletedEvent) EventType() string {
	return CustomerDeletedEventType
}

func (e *CustomerDeletedEvent) OccurredAt() time.Time {
	return e.DeletedAt
}

func (e *OrderCreatedEvent) EventType() string {
	return OrderCreatedEventType
}

func (e *OrderCreatedEvent) OccurredAt() time.Time {
	return e.CreatedAt
}

func (e *OrderUpdatedEvent) EventType() string {
	return OrderUpdatedEventType
}

func (e *OrderUpdatedEvent) OccurredAt() time.Time {
	return e.UpdatedAt
}

func (e *OrderDeletedEvent) EventType() string {
	return OrderDeletedEventType
}

func (e *OrderDeletedEvent) OccurredAt() time.Time {
	return e.DeletedAt
}

func (e *CustomerAssignedToOrderEvent) EventType() string {
	return CustomerAssignedToOrderEventType
}

func (e *CustomerAssignedToOrderEvent) OccurredAt() time.Time {
	return e.AssignedAt
}
