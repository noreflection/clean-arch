package app

import "context"

//type Service interface {
//	Create(ctx context.Context, id int, product string, quantity int) (int, error)
//	Delete(ctx context.Context, id int) error
//	Update(ctx context.Context, id int, quantity int, product string) error
//}

type Service[E any] interface {
	Create(ctx context.Context, id int, entity E) (int, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, entity E) error
}
