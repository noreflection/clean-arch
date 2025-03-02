package usecases

import (
	"context"
	"errors"
	"go-cqrs/internal/app"
	"go-cqrs/internal/interface/dto"
)

// OrderUseCases implements the OrderUseCase interface
type OrderUseCases struct {
	orderRepo    app.OrderRepository
	customerRepo app.CustomerRepository
}

// NewOrderUseCases creates a new OrderUseCases
func NewOrderUseCases(orderRepo app.OrderRepository, customerRepo app.CustomerRepository) *OrderUseCases {
	return &OrderUseCases{
		orderRepo:    orderRepo,
		customerRepo: customerRepo,
	}
}

// CreateOrder implements the OrderUseCase interface
func (s *OrderUseCases) CreateOrder(ctx context.Context, request dto.CreateOrderRequest) (*dto.OrderDTO, error) {
	// Convert DTO to domain entity
	order := request.ToDomain()

	// Validate customer exists if customerID is provided
	if order.CustomerID != nil {
		customer, err := s.customerRepo.GetByID(ctx, *order.CustomerID)
		if err != nil {
			return nil, errors.New("failed to verify customer existence")
		}
		if customer == nil {
			return nil, errors.New("customer not found")
		}
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
func (s *OrderUseCases) GetOrder(ctx context.Context, id int) (*dto.OrderDTO, error) {
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
func (s *OrderUseCases) UpdateOrder(ctx context.Context, request dto.UpdateOrderRequest) error {
	// Check if order exists
	existingOrder, err := s.orderRepo.GetByID(ctx, request.ID)
	if err != nil {
		return err
	}
	if existingOrder == nil {
		return errors.New("order not found")
	}

	// Convert DTO to domain entity
	order := request.ToDomain()

	// Update in repository
	return s.orderRepo.Update(ctx, *order)
}

// DeleteOrder implements the OrderUseCase interface
func (s *OrderUseCases) DeleteOrder(ctx context.Context, id int) error {
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
func (s *OrderUseCases) ListOrders(ctx context.Context, limit, offset int) ([]dto.OrderDTO, error) {
	orders, err := s.orderRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	// Convert domain entities to DTOs
	var orderDTOs []dto.OrderDTO
	for _, order := range orders {
		orderDTOs = append(orderDTOs, dto.ToOrderDTO(order))
	}

	return orderDTOs, nil
}

// AssignCustomerToOrder implements the OrderUseCase interface
func (s *OrderUseCases) AssignCustomerToOrder(ctx context.Context, orderID, customerID int) error {
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
