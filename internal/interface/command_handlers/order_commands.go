package command_handlers

import (
	"context"
	"errors"
	"fmt"
	"go-cqrs/internal/app"
	"go-cqrs/internal/domain"
	"go-cqrs/internal/infra/event_store"
	"log"
)

type OrderCommandHandler struct {
	eventStore event_store.EventStore
	service    app.UseCases[domain.Order]
}

func NewOrderCommandHandler(eventStore event_store.EventStore, service app.UseCases[domain.Order]) *OrderCommandHandler {
	return &OrderCommandHandler{eventStore: eventStore, service: service}
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

	var order domain.Order
	id, err := h.service.Create(ctx, cmd.ID, order)
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

	order := domain.Order{ID: cmd.ID, Product: cmd.Product, Quantity: cmd.Quantity}
	err := h.service.Update(ctx, order)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to update order: %s", err))
	}
	return nil
}
