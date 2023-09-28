package customer

import (
	"go-cqrs/internal/domain"
	"gorm.io/gorm"
)

// CustomerService represents the service for orders.
type CustomerService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *CustomerService {
	return &CustomerService{
		db: db,
	}
}

func (s *CustomerService) CreateCustomer(customer domain.Customer) (*domain.Customer, error) {
	// Implement the CreateOrder method with GORM
	result := s.db.Create(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func (s *CustomerService) GetCustomer(customerID string) (*domain.Customer, error) {
	// Implement the GetOrder method with GORM
	var order domain.Customer
	result := s.db.First(&order, "id = ?", customerID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

// Add other order-related service methods here
