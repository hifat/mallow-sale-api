package recipeHelper

import (
	"context"

	recipeModule "github.com/hifat/mallow-sale-api/internal/recipe"
	recipeRepository "github.com/hifat/mallow-sale-api/internal/recipe/repository"
)

//go:generate mockgen -source=./typeHelper.go -destination=./mock/typeHelper.go -package=mockRecipeHelper
type IRecipeTypeHelper interface {
	FindAndGetByCode(ctx context.Context, codes []recipeModule.EnumCodeRecipeType) (func(code recipeModule.EnumCodeRecipeType) *recipeModule.RecipeTypeResponse, error)
}

type recipeTypeHelper struct {
	recipeTypeRepository recipeRepository.TypeRepository
}

func NewRecipeTypeHelper(recipeTypeRepository recipeRepository.TypeRepository) IRecipeTypeHelper {
	return &recipeTypeHelper{
		recipeTypeRepository: recipeTypeRepository,
	}
}

func (h *recipeTypeHelper) FindAndGetByCode(ctx context.Context, codes []recipeModule.EnumCodeRecipeType) (func(code recipeModule.EnumCodeRecipeType) *recipeModule.RecipeTypeResponse, error) {
	recipeTypes, err := h.recipeTypeRepository.FindInCodes(ctx, codes)
	if err != nil {
		return nil, err
	}

	return func(code recipeModule.EnumCodeRecipeType) *recipeModule.RecipeTypeResponse {
		for _, v := range recipeTypes {
			if v.Code == code {
				return &v
			}
		}

		return nil
	}, nil
}
