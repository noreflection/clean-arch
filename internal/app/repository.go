package app

import "context"

//type Repository interface {
//	Create(order domain.Order) (int, error)
//	Get(orderID int) (domain.Order, error)
//	Update(order domain.Order) error
//	Delete(orderID int) error
//}

//type Repository interface {
//	Create(ctx context.Context, entity interface{}) (int, error)
//	Get(ctx context.Context, id int) (interface{}, error)
//	Update(ctx context.Context, entity interface{}) error
//	Delete(ctx context.Context, id int) error
//}

type Repository[E any] interface {
	Create(ctx context.Context, entity E) (int, error)
	Get(ctx context.Context, id int) (E, error)
	Update(ctx context.Context, entity E) error
	Delete(ctx context.Context, id int) error
}
