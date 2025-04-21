package recipeRepository

import (
	"context"

	"github.com/hifat/cost-calculator-api/internal/recipe"
)

type IRecipeRepository interface {
	Create(ctx context.Context, req recipe.RecipeReq) (id string, err error)
	Find(ctx context.Context) ([]recipe.RecipeRes, error)
	FindByID(ctx context.Context, id string) (*recipe.RecipeRes, error)
	Update(ctx context.Context, id string, req recipe.UpdateRecipeReq) error
	Delete(ctx context.Context, id string) error
}
