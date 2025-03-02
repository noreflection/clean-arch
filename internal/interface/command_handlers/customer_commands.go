package command_handlers

import (
	"context"
	"errors"
	"go-cqrs/internal/app"
	"go-cqrs/internal/infra/event_store"
	"go-cqrs/internal/interface/dto"
)

type CustomerCommandHandler struct {
	eventStore event_store.EventStore
	useCase    app.CustomerUseCase
}

func NewCustomerCommandHandler(eventStore event_store.EventStore, useCase app.CustomerUseCase) *CustomerCommandHandler {
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

	// TODO: Store event
	// event := event_store.NewCustomerCreatedEvent(result.ID, result.Name, result.Email)
	// if err := h.eventStore.StoreEvent(ctx, event); err != nil {
	//     return 0, err
	// }

	return result.ID, nil
}

type DeleteCustomerCommand struct {
	ID int
}

func (h *CustomerCommandHandler) HandleDeleteCustomerCommand(ctx context.Context, cmd DeleteCustomerCommand) error {
	if cmd.ID <= 0 {
		return errors.New("valid ID is required")
	}

	return h.useCase.DeleteCustomer(ctx, cmd.ID)
}

type UpdateCustomerCommand struct {
	ID    int
	Name  string
	Email string
}

func (h *CustomerCommandHandler) HandleUpdateCustomerCommand(ctx context.Context, cmd UpdateCustomerCommand) error {
	if cmd.ID <= 0 {
		return errors.New("valid ID is required")
	}

	request := dto.UpdateCustomerRequest{
		ID:    cmd.ID,
		Name:  cmd.Name,
		Email: cmd.Email,
	}

	return h.useCase.UpdateCustomer(ctx, request)
}
