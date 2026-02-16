package recipeModule

import (
	"context"

	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type IService interface {
	Create(ctx context.Context, req *Request) (*handling.ResponseItem[*Request], error)
	Find(ctx context.Context, query *QueryReq) (*handling.ResponseItems[Response], error)
	FindByID(ctx context.Context, id string) (*handling.ResponseItem[*Response], error)
	UpdateByID(ctx context.Context, id string, req *Request) (*handling.ResponseItem[*Request], error)
	DeleteByID(ctx context.Context, id string) (*handling.ResponseItem[*Request], error)
	UpdateNoBatch(ctx context.Context, reqs []UpdateOrderNoRequest) error
}
