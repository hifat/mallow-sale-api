package promotionRepository

import (
	"context"

	promotionModule "github.com/hifat/mallow-sale-api/internal/promotion"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
)

type IRepository interface {
	Create(ctx context.Context, promotion *promotionModule.Request) error
	Find(ctx context.Context, query *utilsModule.QueryReq) ([]promotionModule.Response, error)
	FindByID(ctx context.Context, id string) (*promotionModule.Response, error)
	UpdateByID(ctx context.Context, id string, promotion *promotionModule.Request) error
	DeleteByID(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
}
