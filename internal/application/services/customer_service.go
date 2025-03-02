package services

import (
	"context"
	"errors"
	"go-cqrs/internal/adapters/http/dto"
	"go-cqrs/internal/application/ports"
)

// CustomerService implements the CustomerUseCase interface
type CustomerService struct {
	customerRepo ports.CustomerRepository
}

// NewCustomerService creates a new CustomerService
func NewCustomerService(customerRepo ports.CustomerRepository) *CustomerService {
	return &CustomerService{customerRepo: customerRepo}
}

// CreateCustomer implements the CustomerUseCase interface
func (s *CustomerService) CreateCustomer(ctx context.Context, request dto.CreateCustomerRequest) (*dto.CustomerDTO, error) {
	// Convert DTO to domain entity
	customer, err := request.ToDomain()
	if err != nil {
		return nil, err
	}

	// Check if email already exists
	existingCustomer, err := s.customerRepo.GetByEmail(ctx, customer.Email)
	if err != nil && !errors.Is(err, errors.New("not found")) {
		return nil, err
	}
	if existingCustomer != nil {
		return nil, errors.New("customer with this email already exists")
	}

	// Save to repository
	customerID, err := s.customerRepo.Create(ctx, *customer)
	if err != nil {
		return nil, err
	}

	// Retrieve the created customer to return
	createdCustomer, err := s.customerRepo.GetByID(ctx, customerID)
	if err != nil {
		return nil, err
	}

	// Convert to DTO and return
	customerDTO := dto.ToCustomerDTO(*createdCustomer)
	return &customerDTO, nil
}

// GetCustomer implements the CustomerUseCase interface
func (s *CustomerService) GetCustomer(ctx context.Context, id int) (*dto.CustomerDTO, error) {
	customer, err := s.customerRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, errors.New("customer not found")
	}

	customerDTO := dto.ToCustomerDTO(*customer)
	return &customerDTO, nil
}

// UpdateCustomer implements the CustomerUseCase interface
func (s *CustomerService) UpdateCustomer(ctx context.Context, request dto.UpdateCustomerRequest) error {
	// Check if customer exists
	existingCustomer, err := s.customerRepo.GetByID(ctx, request.ID)
	if err != nil {
		return err
	}
	if existingCustomer == nil {
		return errors.New("customer not found")
	}

	// Check if email is being changed and if it's already taken
	if existingCustomer.Email != request.Email {
		customerWithEmail, err := s.customerRepo.GetByEmail(ctx, request.Email)
		if err != nil && !errors.Is(err, errors.New("not found")) {
			return err
		}
		if customerWithEmail != nil && customerWithEmail.ID != request.ID {
			return errors.New("email already in use by another customer")
		}
	}

	// Convert DTO to domain entity and update
	customer, err := request.ToDomain()
	if err != nil {
		return err
	}

	return s.customerRepo.Update(ctx, *customer)
}

// DeleteCustomer implements the CustomerUseCase interface
func (s *CustomerService) DeleteCustomer(ctx context.Context, id int) error {
	// Check if customer exists
	existingCustomer, err := s.customerRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingCustomer == nil {
		return errors.New("customer not found")
	}

	// Delete from repository
	return s.customerRepo.Delete(ctx, id)
}

// ListCustomers implements the CustomerUseCase interface
func (s *CustomerService) ListCustomers(ctx context.Context, limit, offset int) ([]dto.CustomerDTO, error) {
	customers, err := s.customerRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	// Convert to DTOs
	customerDTOs := make([]dto.CustomerDTO, len(customers))
	for i, customer := range customers {
		customerDTOs[i] = dto.ToCustomerDTO(customer)
	}

	return customerDTOs, nil
}
