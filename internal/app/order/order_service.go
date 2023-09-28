package order

import (
	"go-cqrs/internal/domain"
	"gorm.io/gorm"
)

// OrderService represents the service for orders.
type OrderService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *OrderService {
	return &OrderService{
		db: db,
	}
}

func (s *OrderService) CreateOrder(order domain.Order) (*domain.Order, error) {
	// Implement the CreateOrder method with GORM
	result := s.db.Create(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (s *OrderService) GetOrder(orderId string) (*domain.Order, error) {
	// Implement the GetOrder method with GORM
	var order domain.Order
	result := s.db.First(&order, "id = ?", orderId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

// Add other order-related service methods here
