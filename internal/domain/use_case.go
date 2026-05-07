package domain

import "context"

type UseCase[T any, R any] interface {
	Execute(ctx context.Context, input T) (R, error)
}
