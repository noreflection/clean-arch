package services

import (
	"context"
	"fmt"
	"go-cqrs/internal/adapters/http/dto"
	"go-cqrs/internal/application/ports"
	"go-cqrs/internal/domain"
	domainerrors "go-cqrs/internal/domain/errors"
)

type OrderService struct {
	orderRepo    ports.OrderRepository
	customerRepo ports.CustomerRepository
}

func NewOrderService(orderRepo ports.OrderRepository, customerRepo ports.CustomerRepository) *OrderService {
	return &OrderService{
		orderRepo:    orderRepo,
		customerRepo: customerRepo,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, request dto.CreateOrderRequest) (*dto.OrderDTO, error) {
	// Create domain entity
	order, err := domain.NewOrder(request.Product, request.Quantity)
	if err != nil {
		return nil, err
	}

	// Assign customer if provided
	if request.CustomerID != nil {
		// Check if customer exists
		customer, err := s.customerRepo.GetByID(ctx, *request.CustomerID)
		if err != nil {
			return nil, fmt.Errorf("failed to check customer: %w", err)
		}
		if customer == nil {
			return nil, domainerrors.NewNotFoundError("customer", *request.CustomerID)
		}

		if err := order.AssignCustomer(*request.CustomerID); err != nil {
			return nil, err
		}
	}

	// Save to repository
	orderID, err := s.orderRepo.Create(ctx, *order)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	// Retrieve the created order to return
	createdOrder, err := s.orderRepo.GetByID(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("order created but failed to retrieve: %w", err)
	}

	orderDTO := dto.ToOrderDTO(*createdOrder)
	return &orderDTO, nil
}

func (s *OrderService) GetOrder(ctx context.Context, id int) (*dto.OrderDTO, error) {
	order, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get order: %w", err)
	}
	if order == nil {
		return nil, domainerrors.NewNotFoundError("order", id)
	}

	orderDTO := dto.ToOrderDTO(*order)
	return &orderDTO, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, request dto.UpdateOrderRequest) error {
	// Check if order exists
	existingOrder, err := s.orderRepo.GetByID(ctx, request.ID)
	if err != nil {
		return fmt.Errorf("failed to find order: %w", err)
	}
	if existingOrder == nil {
		return domainerrors.NewNotFoundError("order", request.ID)
	}

	// Create domain entity with updated values
	updatedOrder, err := domain.NewOrder(request.Product, request.Quantity)
	if err != nil {
		return err
	}
	updatedOrder.ID = request.ID

	// Check and assign customer if provided
	if request.CustomerID != nil {
		customer, err := s.customerRepo.GetByID(ctx, *request.CustomerID)
		if err != nil {
			return fmt.Errorf("failed to check customer: %w", err)
		}
		if customer == nil {
			return domainerrors.NewNotFoundError("customer", *request.CustomerID)
		}

		if err := updatedOrder.AssignCustomer(*request.CustomerID); err != nil {
			return err
		}
	} else if existingOrder.CustomerID != nil {
		// Keep existing customer if not provided
		if err := updatedOrder.AssignCustomer(*existingOrder.CustomerID); err != nil {
			return err
		}
	}

	// Update in repository
	return s.orderRepo.Update(ctx, *updatedOrder)
}

func (s *OrderService) DeleteOrder(ctx context.Context, id int) error {
	// Check if order exists
	order, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to find order: %w", err)
	}
	if order == nil {
		return domainerrors.NewNotFoundError("order", id)
	}

	// Delete from repository
	return s.orderRepo.Delete(ctx, id)
}

func (s *OrderService) ListOrders(ctx context.Context, limit, offset int) ([]dto.OrderDTO, error) {
	orders, err := s.orderRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list orders: %w", err)
	}

	// Convert domain entities to DTOs
	result := make([]dto.OrderDTO, len(orders))
	for i, order := range orders {
		result[i] = dto.ToOrderDTO(order)
	}

	return result, nil
}

func (s *OrderService) AssignCustomerToOrder(ctx context.Context, orderID, customerID int) error {
	// Check if order exists
	order, err := s.orderRepo.GetByID(ctx, orderID)
	if err != nil {
		return fmt.Errorf("failed to find order: %w", err)
	}
	if order == nil {
		return domainerrors.NewNotFoundError("order", orderID)
	}

	// Check if customer exists
	customer, err := s.customerRepo.GetByID(ctx, customerID)
	if err != nil {
		return fmt.Errorf("failed to check customer: %w", err)
	}
	if customer == nil {
		return domainerrors.NewNotFoundError("customer", customerID)
	}

	// Assign customer
	if err := order.AssignCustomer(customerID); err != nil {
		return err
	}

	// Update in repository
	return s.orderRepo.Update(ctx, *order)
}
