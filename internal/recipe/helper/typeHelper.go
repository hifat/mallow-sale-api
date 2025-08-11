package recipeHelper

import (
	"context"

	recipeModule "github.com/hifat/mallow-sale-api/internal/recipe"
	recipeRepository "github.com/hifat/mallow-sale-api/internal/recipe/repository"
)

type RecipeTypeHelper interface {
	FindAndGetByCode(ctx context.Context, codes []string) (func(id string) *recipeModule.RecipeTypeResponse, error)
}

type recipeTypeHelper struct {
	recipeTypeRepository recipeRepository.TypeRepository
}

func NewRecipeTypeHelper(recipeTypeRepository recipeRepository.TypeRepository) RecipeTypeHelper {
	return &recipeTypeHelper{
		recipeTypeRepository: recipeTypeRepository,
	}
}

func (h *recipeTypeHelper) FindAndGetByCode(ctx context.Context, codes []string) (func(id string) *recipeModule.RecipeTypeResponse, error) {
	recipeTypes, err := h.recipeTypeRepository.FindInCodes(ctx, codes)
	if err != nil {
		return nil, err
	}

	return func(code string) *recipeModule.RecipeTypeResponse {
		for _, v := range recipeTypes {
			if v.Code == code {
				return &v
			}
		}

		return nil
	}, nil
}
