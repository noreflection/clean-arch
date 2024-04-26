package app

import "context"

type Repository[E any] interface {
	Create(ctx context.Context, entity E) (int, error)
	Get(ctx context.Context, id int) (E, error)
	Update(ctx context.Context, entity E) error
	Delete(ctx context.Context, id int) error
}
