package queries

import (
	"context"
	"errors"
	"go-cqrs/internal/application/ports"
)

type CustomerQueryHandler struct {
	customerRepo ports.CustomerRepository
}

func NewCustomerQueryHandler(customerRepo ports.CustomerRepository) *CustomerQueryHandler {
	return &CustomerQueryHandler{customerRepo: customerRepo}
}

type GetCustomerQuery struct {
	ID int
}

func (h *CustomerQueryHandler) HandleGetCustomerQuery(ctx context.Context, query GetCustomerQuery) (interface{}, error) {
	customer, err := h.customerRepo.GetByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, errors.New("customer not found")
	}
	return customer, nil
}
