package ports

import (
	"context"
	"go-cqrs/internal/adapters/http/dto"
)

// UseCase is the base interface for all use cases
type UseCase interface{}

// CustomerUseCase defines operations for customer business logic
type CustomerUseCase interface {
	UseCase
	CreateCustomer(ctx context.Context, request dto.CreateCustomerRequest) (*dto.CustomerDTO, error)
	GetCustomer(ctx context.Context, id int) (*dto.CustomerDTO, error)
	UpdateCustomer(ctx context.Context, request dto.UpdateCustomerRequest) error
	DeleteCustomer(ctx context.Context, id int) error
	ListCustomers(ctx context.Context, limit, offset int) ([]dto.CustomerDTO, error)
}

// OrderUseCase defines operations for order business logic
type OrderUseCase interface {
	UseCase
	CreateOrder(ctx context.Context, request dto.CreateOrderRequest) (*dto.OrderDTO, error)
	GetOrder(ctx context.Context, id int) (*dto.OrderDTO, error)
	UpdateOrder(ctx context.Context, request dto.UpdateOrderRequest) error
	DeleteOrder(ctx context.Context, id int) error
	ListOrders(ctx context.Context, limit, offset int) ([]dto.OrderDTO, error)
	AssignCustomerToOrder(ctx context.Context, orderID, customerID int) error
}
