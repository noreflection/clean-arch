package command_handlers

import (
	"context"
	"errors"
	"go-cqrs/internal/app"
	"go-cqrs/internal/infra/event_store"
	"go-cqrs/internal/interface/dto"
)

type OrderCommandHandler struct {
	eventStore event_store.EventStore
	useCase    app.OrderUseCase
}

func NewOrderCommandHandler(eventStore event_store.EventStore, useCase app.OrderUseCase) *OrderCommandHandler {
	return &OrderCommandHandler{eventStore: eventStore, useCase: useCase}
}

type CreateOrderCommand struct {
	CustomerID *int
	Product    string
	Quantity   int
}

func (h *OrderCommandHandler) HandleCreateOrderCommand(ctx context.Context, cmd CreateOrderCommand) (int, error) {
	if cmd.Product == "" {
		return 0, errors.New("product is required")
	}
	if cmd.Quantity <= 0 {
		return 0, errors.New("quantity must be greater than zero")
	}

	request := dto.CreateOrderRequest{
		CustomerID: cmd.CustomerID,
		Product:    cmd.Product,
		Quantity:   cmd.Quantity,
	}

	result, err := h.useCase.CreateOrder(ctx, request)
	if err != nil {
		return 0, err
	}

	// TODO: Store event
	// event := event_store.NewOrderCreatedEvent(result.ID, result.Product, result.Quantity)
	// if err := h.eventStore.StoreEvent(ctx, event); err != nil {
	//     return 0, err
	// }

	return result.ID, nil
}

type DeleteOrderCommand struct {
	ID int
}

func (h *OrderCommandHandler) HandleDeleteOrderCommand(ctx context.Context, cmd DeleteOrderCommand) error {
	if cmd.ID <= 0 {
		return errors.New("valid ID is required")
	}

	return h.useCase.DeleteOrder(ctx, cmd.ID)
}

type UpdateOrderCommand struct {
	ID         int
	CustomerID *int
	Product    string
	Quantity   int
}

func (h *OrderCommandHandler) HandleUpdateOrderCommand(ctx context.Context, cmd UpdateOrderCommand) error {
	if cmd.ID <= 0 {
		return errors.New("valid ID is required")
	}

	request := dto.UpdateOrderRequest{
		ID:         cmd.ID,
		CustomerID: cmd.CustomerID,
		Product:    cmd.Product,
		Quantity:   cmd.Quantity,
	}

	return h.useCase.UpdateOrder(ctx, request)
}

type AssignCustomerCommand struct {
	OrderID    int
	CustomerID int
}

func (h *OrderCommandHandler) HandleAssignCustomerCommand(ctx context.Context, cmd AssignCustomerCommand) error {
	if cmd.OrderID <= 0 {
		return errors.New("valid order ID is required")
	}
	if cmd.CustomerID <= 0 {
		return errors.New("valid customer ID is required")
	}

	return h.useCase.AssignCustomerToOrder(ctx, cmd.OrderID, cmd.CustomerID)
}
