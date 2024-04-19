package command_handlers

import (
	"context"
	"errors"
	"fmt"
	"go-cqrs/internal/app"
	"go-cqrs/internal/domain"
	"go-cqrs/internal/infrastructure/event_store"
	"log"
)

type CustomerCommandHandler struct {
	eventStore event_store.EventStore
	service    app.Service[domain.Customer]
}

func NewCustomerCommandHandler(eventStore event_store.EventStore, service app.Service[domain.Customer]) *CustomerCommandHandler {
	return &CustomerCommandHandler{eventStore: eventStore, service: service}
}

type CreateCustomerCommand struct {
	ID       int
	Product  string
	Quantity int
}

func (h *CustomerCommandHandler) HandleCreateCustomerCommand(ctx context.Context, cmd CreateCustomerCommand) (int, error) {
	if cmd.Product == "" {
		return 0, errors.New("product is required")
	}
	if cmd.Quantity <= 0 {
		return 0, errors.New("quantity must be greater than zero")
	}

	var customer domain.Customer
	id, err := h.service.Create(ctx, cmd.ID, customer)
	if err != nil {
		log.Fatalf("Unable to create order with id %d: %v", id, err)
	}
	//todo: Persist the order creation event
	//event := event_store.NewOrderCreatedEvent(order.ID(), order.Product(), order.Quantity())
	//if err := h.eventStore.StoreEvent(ctx, event); err != nil {
	//	return err
	//}
	return id, nil
}

type DeleteCustomerCommand struct {
	ID int
}

func (h *CustomerCommandHandler) HandleDeleteCustomerCommand(ctx context.Context, cmd DeleteCustomerCommand) error {
	if cmd.ID == 0 {
		return errors.New("ID is required")
	}

	if err := h.service.Delete(ctx, cmd.ID); err != nil {
		return errors.New(fmt.Sprintf("failed to delete customer: %s", err))
	}
	return nil
}

type UpdateCustomerCommand struct {
	ID       int
	Product  string
	Quantity int
}

func (h *CustomerCommandHandler) HandleUpdateCustomerCommand(ctx context.Context, cmd UpdateCustomerCommand) error {
	if cmd.ID == 0 {
		return errors.New("ID is required")
	}

	var customer domain.Customer
	err := h.service.Update(ctx, customer)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to update customer: %s", err))
	}
	return nil
}
