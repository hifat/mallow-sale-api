package purchasePresetModule

import (
	"context"

	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type IService interface {
	Create(ctx context.Context, req *Request) (*handling.ResponseItem[*Response], error)
	Find(ctx context.Context, query *utilsModule.QueryReq) (*handling.ResponseItems[*Response], error)
	FindByID(ctx context.Context, id string) (*handling.ResponseItem[*Response], error)
	UpdateByID(ctx context.Context, id string, req *Request) (*handling.ResponseItem[*Response], error)
	DeleteByID(ctx context.Context, id string) error
}

type IRepository interface {
	Create(ctx context.Context, req *Request) (string, error)
	Find(ctx context.Context, query *utilsModule.QueryReq) ([]Response, error)
	FindByID(ctx context.Context, id string) (*Response, error)
	UpdateByID(ctx context.Context, id string, req *Request) error
	DeleteByID(ctx context.Context, id string) error
	Count(ctx context.Context, query *utilsModule.QueryReq) (int64, error)
}
