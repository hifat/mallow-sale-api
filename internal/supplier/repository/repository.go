package supplierRepository

import (
	"context"

	supplierModule "github.com/hifat/mallow-sale-api/internal/supplier"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
)

type Repository interface {
	Create(ctx context.Context, req *supplierModule.Request) error
	Find(ctx context.Context, query *utilsModule.QueryReq) ([]supplierModule.Response, error)
	FindByID(ctx context.Context, id string) (*supplierModule.Response, error)
	UpdateByID(ctx context.Context, id string, req *supplierModule.Request) error
	DeleteByID(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
	FindInIDs(ctx context.Context, ids []string) ([]supplierModule.Response, error)
}
