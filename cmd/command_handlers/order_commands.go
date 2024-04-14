package command_handlers

import (
	"context"
	"errors"
	"go-cqrs/internal/app/order/repo"
	"go-cqrs/internal/domain"
	"go-cqrs/internal/infrastructure/event_store"
	"log"
)

type OrderCommandHandler struct {
	eventStore event_store.EventStore
	repo       repo.OrderRepository
}

func NewOrderCommandHandler(eventStore event_store.EventStore, repo repo.OrderRepository) *OrderCommandHandler {
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
	order, _ := domain.NewOrder(cmd.ID, cmd.Product, cmd.Quantity)
	id, err := h.repo.Create(*order)
	if err != nil {
		log.Fatalf("Unable to create order with id %d: %v", id, err)
	}
	// Persist the order creation event
	//event := event_store.NewOrderCreatedEvent(order.ID(), order.Product(), order.Quantity())
	//if err := h.eventStore.StoreEvent(ctx, event); err != nil {
	//	return err
	//}

	// Perform any additional logic or validations here

	return nil
}
