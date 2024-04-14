package query_handlers

import (
	"context"
	"fmt"
	"go-cqrs/internal/app/order/repo"
	"go-cqrs/internal/domain"
)

// OrderQueryHandler handles queries related to orders.
type OrderQueryHandler struct {
	orderRepo repo.OrderRepository
}

// NewOrderQueryHandler creates a new instance of OrderQueryHandler.
func NewOrderQueryHandler(orderRepo repo.OrderRepository) *OrderQueryHandler {
	return &OrderQueryHandler{orderRepo: orderRepo}
}

type GetOrderQuery struct {
	ID int
}

func (qh *OrderQueryHandler) HandleGetOrderQuery(ctx context.Context, query GetOrderQuery) (domain.Order, error) {
	order, err := qh.orderRepo.Get(query.ID)
	if err != nil {
		return domain.Order{}, fmt.Errorf("failed to retrieve order by ID: %w", err)
	}
	return order, nil
}
