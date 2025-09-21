package stockRepository

import (
	"context"

	stockModule "github.com/hifat/mallow-sale-api/internal/stock"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
)

type IRepository interface {
	Create(ctx context.Context, req *stockModule.Request) error
	Find(ctx context.Context, query *utilsModule.QueryReq) ([]stockModule.Response, error)
	FindByID(ctx context.Context, id string) (*stockModule.Response, error)
	UpdateByID(ctx context.Context, id string, req *stockModule.Request) error
	DeleteByID(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
}
