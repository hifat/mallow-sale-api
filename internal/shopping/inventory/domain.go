package shoppingInventoryModule

import (
	"context"

	"github.com/hifat/mallow-sale-api/pkg/handling"
)

// Repository
type IRepository interface {
	Create(ctx context.Context, req *Request) error
	Find(ctx context.Context) ([]Response, error)
	DeleteByID(ctx context.Context, id string) error
}

// Service
type IService interface {
	Create(ctx context.Context, req *Request) (*handling.ResponseItem[*Request], error)
	Find(ctx context.Context) (*handling.ResponseItems[Response], error)
	DeleteByID(ctx context.Context, id string) error
}
