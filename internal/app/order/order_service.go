package order

import (
	"database/sql"
	"go-cqrs/internal/domain"
	//"gorm.io/gorm"
)

// OrderService represents the service for orders.
type OrderService struct {
	db *sql.DB
}

func NewService(db *sql.DB) *OrderService {
	return &OrderService{
		db: db,
	}
}

func (s *OrderService) CreateOrder(order domain.Order) (*domain.Order, error) {
	// Implement the CreateOrder method with GORM
	//result := s.db.Create(&order)
	//if result.Error != nil {
	//	return nil, result.Error
	//}
	//
	//return &order, nil
	return nil, nil
}

func (s *OrderService) GetOrder(orderId string) (*domain.Order, error) {
	// Implement the GetOrder method with GORM
	//var order domain.Order
	//result := s.db.First(&order, "id = ?", orderId)
	//if result.Error != nil {
	//	return nil, result.Error
	//}
	//
	//return &order, nil
	return nil, nil
}

// Add other order-related service methods here
