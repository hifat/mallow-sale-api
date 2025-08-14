package inventoryRepository

import (
	"context"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
)

type Repository interface {
	Create(ctx context.Context, req *inventoryModule.Request) error
	Find(ctx context.Context, query *utilsModule.QueryReq) ([]inventoryModule.Response, error)
	FindByID(ctx context.Context, id string) (*inventoryModule.Response, error)
	FindByName(ctx context.Context, name string) (*inventoryModule.Response, error)
	FindInIDs(ctx context.Context, ids []string) ([]inventoryModule.Response, error)
	UpdateByID(ctx context.Context, id string, req *inventoryModule.Request) error
	DeleteByID(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
	UpdateStock(ctx context.Context, id string, quantity float64, purchasePrice float64) error
}
