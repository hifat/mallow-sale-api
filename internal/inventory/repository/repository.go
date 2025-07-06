package inventoryRepository

import (
	"context"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
)

type Repository interface {
	Create(ctx context.Context, req *inventoryModule.Request) error
	Find(ctx context.Context) ([]inventoryModule.Response, error)
	FindByID(ctx context.Context, id string) (*inventoryModule.Response, error)
	FindInIDs(ctx context.Context, ids []string) ([]inventoryModule.Response, error)
	UpdateByID(ctx context.Context, id string, req *inventoryModule.Request) error
	DeleteByID(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
}
