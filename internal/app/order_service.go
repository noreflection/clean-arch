package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"go-cqrs/internal/domain"
	"log"
)

type OrderService struct {
	orderRepo domain.Repository[domain.Order]
}

func NewOrderService(orderRepo domain.Repository[domain.Order]) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (s *OrderService) Create(ctx context.Context, id int, order domain.Order) (int, error) {
	// Additional validation or business logic should be performed here
	// For example, checking if the product exists or if the user is allowed to create orders
	o, _ := domain.NewOrder(id, order.Product, order.Quantity) //todo
	orderID, err := s.orderRepo.Create(ctx, *o)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create order")
	}

	id, err = s.orderRepo.Create(ctx, *o)
	if err != nil {
		log.Fatalf("Unable to create order with id %d: %v", id, err)
	}
	// Any additional logic after creating the order can be added here
	return orderID, nil

}

func (s *OrderService) Delete(ctx context.Context, id int) error {
	if err := s.checkOrderExists(ctx, id); err != nil {
		return err
	}
	err := s.orderRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return err
}

func (s *OrderService) Update(ctx context.Context, order domain.Order) error {
	if err := s.checkOrderExists(ctx, order.ID); err != nil {
		return err
	}

	err := s.orderRepo.Update(ctx, order)
	if err != nil {
		return err
	}
	return err
}

func (s *OrderService) checkOrderExists(ctx context.Context, orderID int) error {
	_, err := s.orderRepo.Get(ctx, orderID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("order with ID %d not found", orderID)
		}
		return errors.Wrap(err, "failed to check if order exists")
	}
	return nil
}