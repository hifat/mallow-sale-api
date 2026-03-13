package middlewareModule

import "context"

type IService interface {
	AuthGuard(ctx context.Context, t string) error
}
