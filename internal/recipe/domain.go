package recipeModule

import (
	"context"

	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

//go:generate mockgen -source=./repository.go -destination=./repository/mock/repository.go -package=mockRecipeRepository
type IRepository interface {
	Create(ctx context.Context, recipe *Request) error
	Find(ctx context.Context, query *QueryReq) ([]Response, error)
	FindInIDs(ctx context.Context, ids []string) ([]Response, error)
	FindByID(ctx context.Context, id string) (*Response, error)
	UpdateByID(ctx context.Context, id string, recipe *Request) error
	DeleteByID(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
	UpdateNoBatch(ctx context.Context, reqs []UpdateOrderNoRequest) error
}

type IService interface {
	Create(ctx context.Context, req *Request) (*handling.ResponseItem[*Request], error)
	Find(ctx context.Context, query *QueryReq) (*handling.ResponseItems[Response], error)
	FindByID(ctx context.Context, id string) (*handling.ResponseItem[*Response], error)
	UpdateByID(ctx context.Context, id string, req *Request) (*handling.ResponseItem[*Request], error)
	DeleteByID(ctx context.Context, id string) (*handling.ResponseItem[*Request], error)
	UpdateNoBatch(ctx context.Context, reqs []UpdateOrderNoRequest) error
}

type TypeRepository interface {
	Find(ctx context.Context, query *utilsModule.QueryReq) ([]RecipeTypeResponse, error)
	FindByCode(ctx context.Context, code EnumCodeRecipeType) (*RecipeTypeResponse, error)
	FindInCodes(ctx context.Context, codes []EnumCodeRecipeType) ([]RecipeTypeResponse, error)
}
