package dto

import (
	"go-cqrs/internal/domain"
)

// CustomerDTO represents the data transfer object for Customer
type CustomerDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateCustomerRequest represents a request to create a customer
type CreateCustomerRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UpdateCustomerRequest represents a request to update a customer
type UpdateCustomerRequest struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ToCustomerDTO converts a domain Customer to a CustomerDTO
func ToCustomerDTO(customer domain.Customer) CustomerDTO {
	return CustomerDTO{
		ID:    customer.ID,
		Name:  customer.Name,
		Email: customer.Email,
	}
}

// ToDomain converts a CustomerDTO to a domain Customer
func (dto CreateCustomerRequest) ToDomain() (*domain.Customer, error) {
	return domain.NewCustomer(dto.Name, dto.Email)
}

// ToDomain converts an UpdateCustomerRequest to a domain Customer
func (dto UpdateCustomerRequest) ToDomain() (*domain.Customer, error) {
	customer, err := domain.NewCustomer(dto.Name, dto.Email)
	if err != nil {
		return nil, err
	}
	customer.ID = dto.ID
	return customer, nil
}
