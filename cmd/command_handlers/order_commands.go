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

func (h *OrderCommandHandler) HandleCreateOrderCommand(ctx context.Context, cmd CreateOrderCommand) error {
	if cmd.Product == "" || cmd.Quantity <= 0 {
		return errors.New("product, and valid quantity are required")
	}

	// Create a new order entity
	order, _ := domain.NewOrder(cmd.ID, cmd.Product, cmd.Quantity)
	id, err := h.repo.Create(*order)
	if err != nil {
		log.Fatalf("Unable to create order with id %d: %v", id, err)
	}
	//todo: Persist the order creation event
	//event := event_store.NewOrderCreatedEvent(order.ID(), order.Product(), order.Quantity())
	//if err := h.eventStore.StoreEvent(ctx, event); err != nil {
	//	return err
	//}

	return nil
}

type DeleteOrderCommand struct {
	ID int
}

func (h *OrderCommandHandler) HandleDeleteOrderCommand(ctx context.Context, cmd DeleteOrderCommand) error {
	if cmd.ID == 0 {
		return errors.New("ID is required")
	}

	if err := h.repo.Delete(cmd.ID); err != nil {
		return errors.New("failed to delete order")
	}
	return nil
}

type UpdateOrderCommand struct {
	ID       string
	Product  string
	Quantity int
}

func (h *OrderCommandHandler) HandleUpdateOrderCommand(ctx context.Context, cmd UpdateOrderCommand) error {
	// Check if the order ID is provided
	if cmd.ID == "" {
		return errors.New("ID is required")
	}

	order, _ := domain.NewOrder(cmd.ID, cmd.Product, cmd.Quantity)
	// Perform validation and update the order in the repository
	err := h.repo.Update(*order)
	if err != nil {
		return errors.New("failed to update order")
	}
	return nil
}
