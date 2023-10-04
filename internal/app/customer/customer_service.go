package customer

import (
	"database/sql"
	"go-cqrs/internal/domain"
	//"gorm.io/gorm"
)

// CustomerService represents the service for orders.
type CustomerService struct {
	db *sql.DB
}

func NewService(db *sql.DB) *CustomerService {
	return &CustomerService{
		db: db,
	}
}

func (s *CustomerService) CreateCustomer(name, email string) (int64, error) {
	cr := NewCustomerRepository(s.db)
	return cr.Create(name, email)
}

func (s *CustomerService) GetCustomer(customerID string) (*domain.Customer, error) {
	// Implement the GetOrder method with GORM
	//var order domain.Customer
	//result := s.db.First(&order, "id = ?", customerID)
	//if result.Error != nil {
	//	return nil, result.Error
	//}

	return nil, nil
}

// Add other order-related service methods here
