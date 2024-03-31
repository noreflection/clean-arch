// order/order_repository_impl.go
package impls

import (
	"github.com/pkg/errors"
	"go-cqrs/internal/domain"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{db: db}
}

func (r *OrderRepositoryImpl) Create(order domain.Order) (int, error) {
	// Insert a new order into the database and return the ID of the newly created order.
	result := r.db.Create(&order)
	if result.Error != nil {
		return 0, errors.Wrap(result.Error, "failed to create order")
	}
	return order.ID, nil
}

func (r *OrderRepositoryImpl) Get(orderID int) (domain.Order, error) {
	// Retrieve an order by ID from the database.
	var order domain.Order
	result := r.db.First(&order, orderID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.Order{}, errors.Wrap(result.Error, "order not found")
		}
		return domain.Order{}, errors.Wrap(result.Error, "failed to get order")
	}
	return order, nil
}

func (r *OrderRepositoryImpl) Update(order domain.Order) error {
	// Update an existing order in the database.
	result := r.db.Save(&order)
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to update order")
	}
	if result.RowsAffected == 0 {
		return errors.New("order not found")
	}
	return nil
}

func (r *OrderRepositoryImpl) Delete(orderID int) error {
	// Delete an order by ID from the database.
	result := r.db.Delete(&domain.Order{}, orderID)
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to delete order")
	}
	if result.RowsAffected == 0 {
		return errors.New("order not found")
	}
	return nil
}

// Implement other interface methods...
