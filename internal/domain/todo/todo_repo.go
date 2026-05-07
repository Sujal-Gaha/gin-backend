package todo

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, todo *Todo) error
	FindByID(ctx context.Context, id string) (*Todo, error)
	FindAll(ctx context.Context) ([]*Todo, error)
	FindByUserID(ctx context.Context, userID string) ([]*Todo, error)
	Update(ctx context.Context, todo *Todo) error
	Delete(ctx context.Context, id string) error
}
