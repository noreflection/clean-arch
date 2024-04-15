package command_handlers

import (
	"context"
	"errors"
	"fmt"
	"go-cqrs/internal/app/order/service"

	"go-cqrs/internal/app/order/repo"
	"go-cqrs/internal/infrastructure/event_store"
	"log"
)

type OrderCommandHandler struct {
	eventStore event_store.EventStore
	repo       repo.OrderRepository
	service    service.OrderService
}

func NewOrderCommandHandler(eventStore event_store.EventStore, repo repo.OrderRepository, service service.OrderService) *OrderCommandHandler {
	return &OrderCommandHandler{eventStore: eventStore, repo: repo, service: service}
}

type CreateOrderCommand struct {
	ID       int
	Product  string
	Quantity int
}

func (h *OrderCommandHandler) HandleCreateOrderCommand(ctx context.Context, cmd CreateOrderCommand) (int, error) {
	if cmd.Product == "" {
		return 0, errors.New("product is required")
	}
	if cmd.Quantity <= 0 {
		return 0, errors.New("quantity must be greater than zero")
	}

	id, err := h.service.Create(ctx, cmd.ID, cmd.Product, cmd.Quantity)
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

type DeleteOrderCommand struct {
	ID int
}

func (h *OrderCommandHandler) HandleDeleteOrderCommand(ctx context.Context, cmd DeleteOrderCommand) error {
	if cmd.ID == 0 {
		return errors.New("ID is required")
	}

	if err := h.service.Delete(ctx, cmd.ID); err != nil {
		return errors.New(fmt.Sprintf("failed to delete order: %s", err))
	}
	return nil
}

type UpdateOrderCommand struct {
	ID       int
	Product  string
	Quantity int
}

func (h *OrderCommandHandler) HandleUpdateOrderCommand(ctx context.Context, cmd UpdateOrderCommand) error {
	if cmd.ID == 0 {
		return errors.New("ID is required")
	}

	//order, _ := domain.NewOrder(cmd.ID, cmd.Product, cmd.Quantity)
	err := h.service.Update(ctx, cmd.ID, cmd.Quantity, cmd.Product)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to update order: %s", err))
	}
	return nil
}
