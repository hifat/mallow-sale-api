package purchaseModule

import (
	"context"

	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type IRepository interface {
	Create(ctx context.Context, req *CreatePurchaseRequest) (string, error)
	Count(ctx context.Context) (int64, error)
	Find(ctx context.Context, query *utilsModule.QueryReq) ([]Response, error)
	FindByID(ctx context.Context, id string) (*Response, error)
	DeleteByID(ctx context.Context, id string) error
	UpdateByID(ctx context.Context, id string, req *CreatePurchaseRequest) error
}

type IService interface {
	Create(ctx context.Context, req *CreatePurchaseRequest) (*handling.ResponseItem[*CreatePurchaseRequest], error)
	Find(ctx context.Context, query *utilsModule.QueryReq) (*handling.ResponseItems[Response], error)
	FindByID(ctx context.Context, id string) (*handling.ResponseItem[*Response], error)
	UpdateByID(ctx context.Context, id string, req *CreatePurchaseRequest) (*handling.ResponseItem[*CreatePurchaseRequest], error)
	DeleteByID(ctx context.Context, id string) error
}
