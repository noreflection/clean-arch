package command_handlers

import (
	"context"
	"errors"
	"go-cqrs/internal/app/impls"
	"go-cqrs/internal/infrastructure/event_store"
)

type OrderCommandHandler struct {
	eventStore event_store.EventStore
	repo       impls.OrderRepositoryImpl
}

func NewOrderCommandHandler(eventStore event_store.EventStore, repo impls.OrderRepositoryImpl) *OrderCommandHandler {
	return &OrderCommandHandler{eventStore: eventStore, repo: repo}
}

type CreateOrderCommand struct {
	ID       string
	Product  string
	Quantity int
}

// HandleCreateOrderCommand handles the CreateOrderCommand.
func (h *OrderCommandHandler) HandleCreateOrderCommand(ctx context.Context, cmd CreateOrderCommand) error {
	// Validate input
	if cmd.ID == "" || cmd.Product == "" || cmd.Quantity <= 0 {
		return errors.New("order ID, product, and valid quantity are required")
	}

	// Create a new order entity
	//o, _ := domain.NewOrder(cmd.ID, cmd.Product, cmd.Quantity)
	//r, err = order.NewOrderRepository()
	// Persist the order creation event
	//event := event_store.NewOrderCreatedEvent(order.ID(), order.Product(), order.Quantity())
	//if err := h.eventStore.StoreEvent(ctx, event); err != nil {
	//	return err
	//}

	// Perform any additional logic or validations here

	return nil
}
