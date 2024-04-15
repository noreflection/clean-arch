package service

import (
	"context"
	"github.com/pkg/errors"
	"go-cqrs/internal/app/order/repo"
	"go-cqrs/internal/domain"
	"log"
)

type OrderService struct {
	orderRepo repo.OrderRepository
}

func NewOrderService(orderRepo repo.OrderRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (s *OrderService) Create(ctx context.Context, id int, product string, quantity int) (int, error) {
	// Additional validation or business logic should be performed here
	// For example, checking if the product exists or if the user is allowed to create orders
	o, _ := domain.NewOrder(id, product, quantity)
	orderID, err := s.orderRepo.Create(*o)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create order")
	}

	id, err = s.orderRepo.Create(*o)
	if err != nil {
		log.Fatalf("Unable to create order with id %d: %v", id, err)
	}
	// Any additional logic after creating the order can be added here
	return orderID, nil

}

func (s *OrderService) Delete(ctx context.Context, id int) error {
	err := s.orderRepo.Delete(id)
	if err != nil {
		return err
	}
	return err
}
