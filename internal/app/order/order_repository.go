package order

import (
	"github.com/pkg/errors"
	"go-cqrs/internal/domain"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Create(order domain.Order) (int, error) {
	// Insert a new order into the database and return the ID of the newly created order.
	result := r.db.Create(&order)
	if result.Error != nil {
		return 0, errors.Wrap(result.Error, "failed to create order")
	}
	return order.ID, nil
}

func (r *OrderRepository) Get(orderID int) (domain.Order, error) {
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

func (r *OrderRepository) Update(order domain.Order) error {
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

func (r *OrderRepository) Delete(orderID int) error {
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
