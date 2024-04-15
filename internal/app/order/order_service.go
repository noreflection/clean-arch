package order

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"go-cqrs/internal/app"
	"go-cqrs/internal/domain"
	"log"
)

type OrderService struct {
	orderRepo app.Repository
}

func NewOrderService(orderRepo app.Repository) *OrderService {
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
	if err := s.checkOrderExists(id); err != nil {
		return err
	}
	err := s.orderRepo.Delete(id)
	if err != nil {
		return err
	}
	return err
}

func (s *OrderService) Update(ctx context.Context, id int, quantity int, product string) error {
	if err := s.checkOrderExists(id); err != nil {
		return err
	}

	order, _ := domain.NewOrder(id, product, quantity)
	err := s.orderRepo.Update(*order)
	if err != nil {
		return err
	}
	return err
}

func (s *OrderService) checkOrderExists(orderID int) error {
	_, err := s.orderRepo.Get(orderID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("order with ID %d not found", orderID)
		}
		return errors.Wrap(err, "failed to check if order exists")
	}
	return nil
}
