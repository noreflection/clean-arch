package query_handlers

import (
	"context"
	"fmt"
	"go-cqrs/internal/app"
	"go-cqrs/internal/domain"
)

type OrderQueryHandler struct {
	orderRepo app.Repository[domain.Order]
}

func NewOrderQueryHandler(orderRepo app.Repository[domain.Order]) *OrderQueryHandler {
	return &OrderQueryHandler{orderRepo: orderRepo}
}

type GetOrderQuery struct {
	ID int
}

func (h *OrderQueryHandler) HandleGetOrderQuery(ctx context.Context, query GetOrderQuery) (interface{}, error) {
	order, err := h.orderRepo.Get(ctx, query.ID)
	if err != nil {
		return domain.Order{}, fmt.Errorf("failed to retrieve order by ID: %w", err)
	}
	return order, nil
}
