package ports

import (
	"context"
	"go-cqrs/internal/domain"
)

// Repository is the base interface for all repositories
type Repository interface{}

// CustomerRepository defines operations for customer persistence
type CustomerRepository interface {
	Repository
	Create(ctx context.Context, customer domain.Customer) (int, error)
	GetByID(ctx context.Context, id int) (*domain.Customer, error)
	GetByEmail(ctx context.Context, email string) (*domain.Customer, error)
	Update(ctx context.Context, customer domain.Customer) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, limit, offset int) ([]domain.Customer, error)
}

// OrderRepository defines operations for order persistence
type OrderRepository interface {
	Repository
	Create(ctx context.Context, order domain.Order) (int, error)
	GetByID(ctx context.Context, id int) (*domain.Order, error)
	GetByCustomerID(ctx context.Context, customerID int) ([]domain.Order, error)
	Update(ctx context.Context, order domain.Order) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, limit, offset int) ([]domain.Order, error)
}
