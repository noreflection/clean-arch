package app

import "context"

type Service[E any] interface {
	Create(ctx context.Context, id int, entity E) (int, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, entity E) error
}
