// order_commands.go

package command_handlers

import (
	"context"
	"errors"
	"go-cqrs/internal/domain"
	"go-cqrs/internal/infrastructure/event_store"
)

// OrderCommandHandler handles order-related commands.
type OrderCommandHandler struct {
	eventStore event_store.OrderEventStore
}

// NewOrderCommandHandler creates a new instance of OrderCommandHandler.
func NewOrderCommandHandler(eventStore event_store.OrderEventStore) *OrderCommandHandler {
	return &OrderCommandHandler{eventStore: eventStore}
}

// CreateOrderCommand is a command to create a new order.
type CreateOrderCommand struct {
	ID       string
	Product  string
	Quantity int
	// Add other order attributes here
}

// HandleCreateOrderCommand handles the CreateOrderCommand.
func (h *OrderCommandHandler) HandleCreateOrderCommand(ctx context.Context, cmd CreateOrderCommand) error {
	// Validate input
	if cmd.ID == "" || cmd.Product == "" || cmd.Quantity <= 0 {
		return errors.New("order ID, product, and valid quantity are required")
	}

	// Create a new order entity
	order := domain.NewOrder(cmd.ID, cmd.Product, cmd.Quantity)

	// Persist the order creation event
	event := event_store.NewOrderCreatedEvent(order.ID(), order.Product(), order.Quantity())
	if err := h.eventStore.StoreEvent(ctx, event); err != nil {
		return err
	}

	// Perform any additional logic or validations here

	return nil
}

// Add other order-related commands and their handlers here
