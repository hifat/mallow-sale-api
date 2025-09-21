package recipeRepository

import (
	"context"

	recipeModule "github.com/hifat/mallow-sale-api/internal/recipe"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
)

type IRepository interface {
	Create(ctx context.Context, recipe *recipeModule.Request) error
	Find(ctx context.Context, query *recipeModule.QueryReq) ([]recipeModule.Response, error)
	FindInIDs(ctx context.Context, ids []string) ([]recipeModule.Response, error)
	FindByID(ctx context.Context, id string) (*recipeModule.Response, error)
	UpdateByID(ctx context.Context, id string, recipe *recipeModule.Request) error
	DeleteByID(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
	UpdateNoBatch(ctx context.Context, reqs []recipeModule.UpdateOrderNoRequest) error
}

type TypeRepository interface {
	Find(ctx context.Context, query *utilsModule.QueryReq) ([]recipeModule.RecipeTypeResponse, error)
	FindByCode(ctx context.Context, code string) (*recipeModule.RecipeTypeResponse, error)
	FindInCodes(ctx context.Context, codes []string) ([]recipeModule.RecipeTypeResponse, error)
}
