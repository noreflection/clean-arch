// customer_commands.go
package command_handlers

import (
	"context"
	"errors"
	//"go-cqrs/internal/app/customer"
	"go-cqrs/internal/domain"
	"gorm.io/gorm"
)

// CustomerCommandHandler handles customer-related commands.
type CustomerCommandHandler struct {
	//eventStore event_store.EventStore
	db      *gorm.DB
	service domain.CustomerService
}

// NewCustomerCommandHandler creates a new instance of CustomerCommandHandler.
func NewCustomerCommandHandler(d *gorm.DB) *CustomerCommandHandler {
	return &CustomerCommandHandler{ /*eventStore: eventStore*/ db: d}
}

// CreateCustomerCommand is a command to create a new customer.
type CreateCustomerCommand struct {
	ID    string
	Name  string
	Email string
}

// HandleCreateCustomerCommand handles the CreateCustomerCommand.
func (h *CustomerCommandHandler) HandleCreateCustomerCommand(ctx context.Context, cmd CreateCustomerCommand) error {
	// Validate input
	if cmd.ID == "" || cmd.Name == "" || cmd.Email == "" {
		return errors.New("customer ID, name, and email are required")
	}

	// Use the customer service to create a new customer
	err := h.service.CreateCustomer(cmd.ID, cmd.Name, cmd.Email)
	if err != nil {
		return err
	}

	return nil

	// Create a new customer entity
	//customer := domain.NewCustomer( /*cmd.ID,*/ cmd.Name, cmd.Email)
	//database := command_database.NewOrderCommandDB()
	// Persist the customer creation
	//
	// event store
	//event := event_store.NewCustomerCreatedEvent(customer.ID, customer.Name, customer.Email)
	//if err := h.eventStore.StoreEvent(ctx, *event); err != nil {
	//	return err
	//}

	// Perform any additional logic or validations here

	//return nil
}

// Add other customer-related commands and their handlers here
