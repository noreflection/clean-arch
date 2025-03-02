package query_handlers

import (
	"context"
	"errors"
	"go-cqrs/internal/app"
)

type OrderQueryHandler struct {
	orderRepo app.OrderRepository
}

func NewOrderQueryHandler(orderRepo app.OrderRepository) *OrderQueryHandler {
	return &OrderQueryHandler{orderRepo: orderRepo}
}

type GetOrderQuery struct {
	ID int
}

func (h *OrderQueryHandler) HandleGetOrderQuery(ctx context.Context, query GetOrderQuery) (interface{}, error) {
	order, err := h.orderRepo.GetByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("order not found")
	}
	return order, nil
}
