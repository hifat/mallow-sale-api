package inventoryModule

import (
	"context"

	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
)

//go:generate mockgen -source=./domain.go -destination=./mock/repository/mock/mongo.go -package=mockInventoryRepository
type IRepository interface {
	Create(ctx context.Context, req *Request) error
	Find(ctx context.Context, query *utilsModule.QueryReq) ([]Response, error)
	FindByID(ctx context.Context, id string) (*Response, error)
	FindByName(ctx context.Context, name string) (*Response, error)
	FindInIDs(ctx context.Context, ids []string) ([]Response, error)
	UpdateByID(ctx context.Context, id string, req *Request) error
	DeleteByID(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
	UpdateStock(ctx context.Context, id string, quantity float64, purchasePrice float64) error
	UpdatePurchasePrice(ctx context.Context, id string, purchasePrice float64) error
}
