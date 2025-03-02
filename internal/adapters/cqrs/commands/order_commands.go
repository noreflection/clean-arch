package commands

import (
	"context"
	"errors"
	"fmt"
	"go-cqrs/internal/adapters/http/dto"
	"go-cqrs/internal/application/ports"
	"go-cqrs/internal/domain/events"
	event_store "go-cqrs/internal/infrastructure/messaging/events"
	"strconv"
)

type OrderCommandHandler struct {
	eventStore event_store.EventStore
	useCase    ports.OrderUseCase
}

func NewOrderCommandHandler(eventStore event_store.EventStore, useCase ports.OrderUseCase) *OrderCommandHandler {
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

	// Store event
	event := events.NewOrderCreatedEvent(
		strconv.Itoa(result.ID),
		result.Product,
		result.Quantity,
	)
	if err := h.eventStore.StoreEvent(ctx, event); err != nil {
		// Log the error but don't fail the operation
		fmt.Printf("Warning: Failed to store order created event: %v\n", err)
	}

	return result.ID, nil
}

type DeleteOrderCommand struct {
	ID int
}

func (h *OrderCommandHandler) HandleDeleteOrderCommand(ctx context.Context, cmd DeleteOrderCommand) error {
	if cmd.ID <= 0 {
		return errors.New("invalid order ID")
	}

	err := h.useCase.DeleteOrder(ctx, cmd.ID)
	if err != nil {
		return err
	}

	// Record the order deleted event
	event := events.NewOrderDeletedEvent(strconv.Itoa(cmd.ID))
	if err := h.eventStore.StoreEvent(ctx, event); err != nil {
		fmt.Printf("Warning: Failed to store order deleted event: %v\n", err)
	}

	return nil
}

type UpdateOrderCommand struct {
	ID         int
	CustomerID *int
	Product    string
	Quantity   int
}

func (h *OrderCommandHandler) HandleUpdateOrderCommand(ctx context.Context, cmd UpdateOrderCommand) error {
	if cmd.ID <= 0 {
		return errors.New("invalid order ID")
	}
	if cmd.Product == "" {
		return errors.New("product is required")
	}
	if cmd.Quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	request := dto.UpdateOrderRequest{
		ID:         cmd.ID,
		CustomerID: cmd.CustomerID,
		Product:    cmd.Product,
		Quantity:   cmd.Quantity,
	}

	err := h.useCase.UpdateOrder(ctx, request)
	if err != nil {
		return err
	}

	// Record the order updated event
	var customerIDStr *string
	if cmd.CustomerID != nil {
		idStr := strconv.Itoa(*cmd.CustomerID)
		customerIDStr = &idStr
	}

	event := events.NewOrderUpdatedEvent(
		strconv.Itoa(cmd.ID),
		cmd.Product,
		cmd.Quantity,
		customerIDStr,
	)
	if err := h.eventStore.StoreEvent(ctx, event); err != nil {
		fmt.Printf("Warning: Failed to store order updated event: %v\n", err)
	}

	return nil
}

type AssignCustomerCommand struct {
	OrderID    int
	CustomerID int
}

func (h *OrderCommandHandler) HandleAssignCustomerCommand(ctx context.Context, cmd AssignCustomerCommand) error {
	if cmd.OrderID <= 0 {
		return errors.New("invalid order ID")
	}
	if cmd.CustomerID <= 0 {
		return errors.New("invalid customer ID")
	}

	err := h.useCase.AssignCustomerToOrder(ctx, cmd.OrderID, cmd.CustomerID)
	if err != nil {
		return err
	}

	// Record the customer assigned event
	event := events.NewCustomerAssignedToOrderEvent(
		strconv.Itoa(cmd.OrderID),
		strconv.Itoa(cmd.CustomerID),
	)
	if err := h.eventStore.StoreEvent(ctx, event); err != nil {
		fmt.Printf("Warning: Failed to store customer assigned event: %v\n", err)
	}

	return nil
}
