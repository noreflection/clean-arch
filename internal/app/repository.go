package app

import "go-cqrs/internal/domain"

type Repository interface {
	Create(order domain.Order) (int, error)
	Get(orderID int) (domain.Order, error)
	Update(order domain.Order) error
	Delete(orderID int) error
}
