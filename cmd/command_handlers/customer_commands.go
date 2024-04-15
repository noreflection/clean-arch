package command_handlers

import (
	"database/sql"
	"errors"
	"go-cqrs/internal/infrastructure/event_store"
)

type CustomerCommandHandler struct {
	//eventStore event_store.EventStore
	db *sql.DB
	//service domain.CustomerService
}

func NewCustomerCommandHandler(store event_store.EventStore, d *sql.DB) *CustomerCommandHandler {
	return &CustomerCommandHandler{ /*eventStore: eventStore*/ db: d}
}

type CreateCustomerCommand struct {
	ID    string
	Name  string
	Email string
}

func (h *CustomerCommandHandler) HandleCreateCustomerCommand(cmd CreateCustomerCommand) error {
	if cmd.ID == "" || cmd.Name == "" || cmd.Email == "" {
		return errors.New("customer ID, name, and email are required")
	}

	//err := h.service.CreateCustomer(cmd.ID, cmd.Name, cmd.Email)
	//if err != nil {
	//	return err
	//}

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
