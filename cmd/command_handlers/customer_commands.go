package command_handlers

import (
	"context"
	"go-cqrs/internal/app/customer"
	"go-cqrs/internal/domain"
)

// CustomerCommandHandler handles customer-related commands.
type CustomerCommandHandler struct {
	customerService customer.CustomerService
}

// NewCustomerCommandHandler creates a new CustomerCommandHandler.
func NewCustomerCommandHandler(customerService customer.CustomerService) *CustomerCommandHandler {
	return &CustomerCommandHandler{
		customerService: customerService,
	}
}

// CreateCustomerCommand represents a command to create a new customer.
type CreateCustomerCommand struct {
	Name  string
	Email string
}

// CreateCustomerCommand handles the creation of a new customer.
func (h *CustomerCommandHandler) CreateCustomerCommand(ctx context.Context, cmd CreateCustomerCommand) (*domain.Customer, error) {
	// Create a new customer entity.
	newCustomer := domain.NewCustomer(cmd.Name, cmd.Email)

	// Use the customer service to create the customer.
	createdCustomer, err := h.customerService.CreateCustomer(*newCustomer)
	if err != nil {
		return nil, err
	}
	// You can add event publishing logic here if needed.
	return createdCustomer, nil
}

// UpdateCustomerCommand handles the updating of an existing customer.
func (h *CustomerCommandHandler) UpdateCustomerCommand(ctx context.Context, customer domain.Customer) (*domain.Customer, error) {
	// Use the customer service to update the customer.
	updatedCustomer, err := h.customerService.UpdateCustomer(customer)
	if err != nil {
		return nil, err
	}
	// You can add event publishing logic here if needed.
	return updatedCustomer, nil
}

// DeleteCustomerCommand handles the deletion of an existing customer.
func (h *CustomerCommandHandler) DeleteCustomerCommand(ctx context.Context, customerID string) error {
	// Use the customer service to delete the customer.
	err := h.customerService.DeleteCustomer(customerID)
	if err != nil {
		return err
	}
	// You can add event publishing logic here if needed.
	return nil
}

// Other customer-related command handlers (e.g., UpdateCustomerCommand, DeleteCustomerCommand) go here.
