package dto

import (
	"go-cqrs/internal/domain"
)

// OrderDTO represents the data transfer object for Order
type OrderDTO struct {
	ID         int    `json:"id"`
	CustomerID *int   `json:"customerId,omitempty"`
	Product    string `json:"product"`
	Quantity   int    `json:"quantity"`
	Status     string `json:"status,omitempty"`
}

// CreateOrderRequest represents a request to create an order
type CreateOrderRequest struct {
	CustomerID *int   `json:"customerId,omitempty"`
	Product    string `json:"product"`
	Quantity   int    `json:"quantity"`
}

// UpdateOrderRequest represents a request to update an order
type UpdateOrderRequest struct {
	ID         int    `json:"id"`
	CustomerID *int   `json:"customerId,omitempty"`
	Product    string `json:"product"`
	Quantity   int    `json:"quantity"`
}

// AssignCustomerRequest represents a request to assign a customer to an order
type AssignCustomerRequest struct {
	OrderID    int `json:"orderId"`
	CustomerID int `json:"customerId"`
}

// ToOrderDTO converts a domain Order to an OrderDTO
func ToOrderDTO(order domain.Order) OrderDTO {
	return OrderDTO{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		Product:    order.Product,
		Quantity:   order.Quantity,
	}
}

// ToDomain converts a CreateOrderRequest to a domain Order
func (dto CreateOrderRequest) ToDomain() (*domain.Order, error) {
	order, err := domain.NewOrder(dto.Product, dto.Quantity)
	if err != nil {
		return nil, err
	}

	if dto.CustomerID != nil {
		if err := order.AssignCustomer(*dto.CustomerID); err != nil {
			return nil, err
		}
	}
	return order, nil
}

// ToDomain converts an UpdateOrderRequest to a domain Order
func (dto UpdateOrderRequest) ToDomain() (*domain.Order, error) {
	order, err := domain.NewOrder(dto.Product, dto.Quantity)
	if err != nil {
		return nil, err
	}

	order.ID = dto.ID
	if dto.CustomerID != nil {
		if err := order.AssignCustomer(*dto.CustomerID); err != nil {
			return nil, err
		}
	}
	return order, nil
}
