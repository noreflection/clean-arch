package query_handlers

import (
	"context"
	"fmt"
	"go-cqrs/internal/domain"
)

type CustomerQueryHandler struct {
	customerRepo domain.Repository[domain.Customer]
}

func NewCustomerQueryHandler(customerRepo domain.Repository[domain.Customer]) *CustomerQueryHandler {
	return &CustomerQueryHandler{customerRepo: customerRepo}
}

type GetCustomerQuery struct {
	ID int
}

func (h *CustomerQueryHandler) HandleGetCustomerQuery(ctx context.Context, query GetCustomerQuery) (interface{}, error) {
	customer, err := h.customerRepo.Get(ctx, query.ID)
	if err != nil {
		return domain.Customer{}, fmt.Errorf("failed to retrieve customer by ID: %w", err)
	}
	return customer, nil
}
