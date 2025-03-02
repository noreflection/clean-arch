package services

import (
	"context"
	"errors"
	"go-cqrs/internal/adapters/http/dto"
	"go-cqrs/internal/application/ports"
	"go-cqrs/internal/domain"
)

// OrderService implements the OrderUseCase interface
type OrderService struct {
	orderRepo    ports.OrderRepository
	customerRepo ports.CustomerRepository
}

// NewOrderService creates a new OrderService
func NewOrderService(orderRepo ports.OrderRepository, customerRepo ports.CustomerRepository) *OrderService {
	return &OrderService{
		orderRepo:    orderRepo,
		customerRepo: customerRepo,
	}
}

// CreateOrder implements the OrderUseCase interface
func (s *OrderService) CreateOrder(ctx context.Context, request dto.CreateOrderRequest) (*dto.OrderDTO, error) {
	// Create domain entity
	order := domain.NewOrder(request.Product, request.Quantity)

	// Assign customer if provided
	if request.CustomerID != nil {
		// Check if customer exists
		customer, err := s.customerRepo.GetByID(ctx, *request.CustomerID)
		if err != nil {
			return nil, err
		}
		if customer == nil {
			return nil, errors.New("customer not found")
		}

		order.AssignCustomer(*request.CustomerID)
	}

	// Save to repository
	orderID, err := s.orderRepo.Create(ctx, *order)
	if err != nil {
		return nil, err
	}

	// Retrieve the created order to return
	createdOrder, err := s.orderRepo.GetByID(ctx, orderID)
	if err != nil {
		return nil, err
	}

	// Convert to DTO and return
	orderDTO := dto.ToOrderDTO(*createdOrder)
	return &orderDTO, nil
}

// GetOrder implements the OrderUseCase interface
func (s *OrderService) GetOrder(ctx context.Context, id int) (*dto.OrderDTO, error) {
	order, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("order not found")
	}

	orderDTO := dto.ToOrderDTO(*order)
	return &orderDTO, nil
}

// UpdateOrder implements the OrderUseCase interface
func (s *OrderService) UpdateOrder(ctx context.Context, request dto.UpdateOrderRequest) error {
	// Check if order exists
	existingOrder, err := s.orderRepo.GetByID(ctx, request.ID)
	if err != nil {
		return err
	}
	if existingOrder == nil {
		return errors.New("order not found")
	}

	// Update order fields
	order := domain.Order{
		ID:       request.ID,
		Product:  request.Product,
		Quantity: request.Quantity,
	}

	// Keep existing customer ID if not changing
	if request.CustomerID != nil {
		order.CustomerID = request.CustomerID
	} else {
		order.CustomerID = existingOrder.CustomerID
	}

	// Update in repository
	return s.orderRepo.Update(ctx, order)
}

// DeleteOrder implements the OrderUseCase interface
func (s *OrderService) DeleteOrder(ctx context.Context, id int) error {
	// Check if order exists
	existingOrder, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingOrder == nil {
		return errors.New("order not found")
	}

	// Delete from repository
	return s.orderRepo.Delete(ctx, id)
}

// ListOrders implements the OrderUseCase interface
func (s *OrderService) ListOrders(ctx context.Context, limit, offset int) ([]dto.OrderDTO, error) {
	orders, err := s.orderRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	// Convert to DTOs
	orderDTOs := make([]dto.OrderDTO, len(orders))
	for i, order := range orders {
		orderDTOs[i] = dto.ToOrderDTO(order)
	}

	return orderDTOs, nil
}

// AssignCustomerToOrder implements the OrderUseCase interface
func (s *OrderService) AssignCustomerToOrder(ctx context.Context, orderID, customerID int) error {
	// Check if order exists
	order, err := s.orderRepo.GetByID(ctx, orderID)
	if err != nil {
		return err
	}
	if order == nil {
		return errors.New("order not found")
	}

	// Check if customer exists
	customer, err := s.customerRepo.GetByID(ctx, customerID)
	if err != nil {
		return err
	}
	if customer == nil {
		return errors.New("customer not found")
	}

	// Assign customer to order
	order.AssignCustomer(customerID)

	// Update in repository
	return s.orderRepo.Update(ctx, *order)
}
