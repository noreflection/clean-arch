package customer

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"go-cqrs/internal/app"
	"go-cqrs/internal/domain"
	"log"
)

type Service struct {
	customerRepo app.Repository[domain.Customer]
}

func NewService(customerRepo app.Repository[domain.Customer]) *Service {
	return &Service{customerRepo: customerRepo}
}

func (s *Service) Create(ctx context.Context, id int, customer domain.Customer) (int, error) {
	// Additional validation or business logic should be performed here
	// For example, checking if the customer exists or if the user is allowed to create orders
	c := domain.NewCustomer(customer.Name, customer.Email)
	orderID, err := s.customerRepo.Create(ctx, *c)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create order")
	}

	id, err = s.customerRepo.Create(ctx, *c)
	if err != nil {
		log.Fatalf("Unable to create order with id %d: %v", id, err)
	}
	return orderID, nil
}

func (s *Service) Delete(ctx context.Context, id int) error {
	if err := s.checkCustomerExists(ctx, id); err != nil {
		return err
	}

	err := s.customerRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) Update(ctx context.Context, customer domain.Customer) error {
	if err := s.checkCustomerExists(ctx, customer.ID); err != nil {
		return err
	}

	err := s.customerRepo.Update(ctx, customer)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) checkCustomerExists(ctx context.Context, orderID int) error {
	_, err := s.customerRepo.Get(ctx, orderID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("customer with ID %d not found", orderID)
		}
		return errors.Wrap(err, "failed to check if customer exists")
	}
	return nil
}
