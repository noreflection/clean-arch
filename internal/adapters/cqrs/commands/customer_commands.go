package commands

import (
	"context"
	"errors"
	"fmt"
	"go-cqrs/internal/adapters/http/dto"
	"go-cqrs/internal/application/ports"
	event_store "go-cqrs/internal/infrastructure/messaging/events"
	"strconv"
)

type CustomerCommandHandler struct {
	eventStore event_store.EventStore
	useCase    ports.CustomerUseCase
}

func NewCustomerCommandHandler(eventStore event_store.EventStore, useCase ports.CustomerUseCase) *CustomerCommandHandler {
	return &CustomerCommandHandler{eventStore: eventStore, useCase: useCase}
}

type CreateCustomerCommand struct {
	Name  string
	Email string
}

func (h *CustomerCommandHandler) HandleCreateCustomerCommand(ctx context.Context, cmd CreateCustomerCommand) (int, error) {
	if cmd.Name == "" {
		return 0, errors.New("name is required")
	}
	if cmd.Email == "" {
		return 0, errors.New("email is required")
	}

	request := dto.CreateCustomerRequest{
		Name:  cmd.Name,
		Email: cmd.Email,
	}

	result, err := h.useCase.CreateCustomer(ctx, request)
	if err != nil {
		return 0, err
	}

	// Store event
	event := event_store.NewCustomerCreatedEvent(
		strconv.Itoa(result.ID),
		result.Name,
		result.Email,
	)
	if err := h.eventStore.StoreEvent(ctx, event); err != nil {
		// Log the error but don't fail the operation
		fmt.Printf("Warning: Failed to store customer created event: %v\n", err)
	}

	return result.ID, nil
}

type DeleteCustomerCommand struct {
	ID int
}

func (h *CustomerCommandHandler) HandleDeleteCustomerCommand(ctx context.Context, cmd DeleteCustomerCommand) error {
	if cmd.ID <= 0 {
		return errors.New("invalid customer ID")
	}

	err := h.useCase.DeleteCustomer(ctx, cmd.ID)
	if err != nil {
		return err
	}

	// Record customer deleted event
	if err := h.eventStore.StoreEvent(ctx, map[string]interface{}{
		"type":       "customer_deleted",
		"customerId": cmd.ID,
	}); err != nil {
		fmt.Printf("Warning: Failed to store customer deleted event: %v\n", err)
	}

	return nil
}

type UpdateCustomerCommand struct {
	ID    int
	Name  string
	Email string
}

func (h *CustomerCommandHandler) HandleUpdateCustomerCommand(ctx context.Context, cmd UpdateCustomerCommand) error {
	if cmd.ID <= 0 {
		return errors.New("invalid customer ID")
	}
	if cmd.Name == "" {
		return errors.New("name is required")
	}
	if cmd.Email == "" {
		return errors.New("email is required")
	}

	request := dto.UpdateCustomerRequest{
		ID:    cmd.ID,
		Name:  cmd.Name,
		Email: cmd.Email,
	}

	err := h.useCase.UpdateCustomer(ctx, request)
	if err != nil {
		return err
	}

	// Record customer updated event
	if err := h.eventStore.StoreEvent(ctx, map[string]interface{}{
		"type":       "customer_updated",
		"customerId": cmd.ID,
		"name":       cmd.Name,
		"email":      cmd.Email,
	}); err != nil {
		fmt.Printf("Warning: Failed to store customer updated event: %v\n", err)
	}

	return nil
}
