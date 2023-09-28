// customer_commands.go

package command_handlers

import (
	"context"
	"errors"
	"go-cqrs/internal/domain"
	"go-cqrs/internal/infrastructure/event_store"
)

// CustomerCommandHandler handles customer-related commands.
type CustomerCommandHandler struct {
	eventStore event_store.CustomerEventStore
}

// NewCustomerCommandHandler creates a new instance of CustomerCommandHandler.
func NewCustomerCommandHandler(eventStore event_store.CustomerEventStore) *CustomerCommandHandler {
	return &CustomerCommandHandler{eventStore: eventStore}
}

// CreateCustomerCommand is a command to create a new customer.
type CreateCustomerCommand struct {
	ID    string
	Name  string
	Email string
	// Add other customer attributes here
}

// HandleCreateCustomerCommand handles the CreateCustomerCommand.
func (h *CustomerCommandHandler) HandleCreateCustomerCommand(ctx context.Context, cmd CreateCustomerCommand) error {
	// Validate input
	if cmd.ID == "" || cmd.Name == "" || cmd.Email == "" {
		return errors.New("customer ID, name, and email are required")
	}

	// Create a new customer entity
	customer := domain.NewCustomer( /*cmd.ID,*/ cmd.Name, cmd.Email)

	// Persist the customer creation event
	event := event_store.NewCustomerCreatedEvent(customer.ID, customer.Name, customer.Email)
	if err := h.eventStore.StoreEvent(ctx, *event); err != nil {
		return err
	}

	// Perform any additional logic or validations here

	return nil
}

// Add other customer-related commands and their handlers here
