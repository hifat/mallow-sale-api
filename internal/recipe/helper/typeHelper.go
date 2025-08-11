package recipeHelper

import (
	"context"

	recipeModule "github.com/hifat/mallow-sale-api/internal/recipe"
	recipeRepository "github.com/hifat/mallow-sale-api/internal/recipe/repository"
)

type TypeHelper interface {
	FindAndGetByCode(ctx context.Context, codes []string) (func(id string) *recipeModule.TypeResponse, error)
}

type typeHelper struct {
	recipeTypeRepository recipeRepository.TypeRepository
}

func NewType(recipeTypeRepository recipeRepository.TypeRepository) TypeHelper {
	return &typeHelper{
		recipeTypeRepository: recipeTypeRepository,
	}
}

func (h *typeHelper) FindAndGetByCode(ctx context.Context, codes []string) (func(id string) *recipeModule.TypeResponse, error) {
	recipeTypes, err := h.recipeTypeRepository.FindInCodes(ctx, codes)
	if err != nil {
		return nil, err
	}

	return func(code string) *recipeModule.TypeResponse {
		for _, v := range recipeTypes {
			if v.Code == code {
				return &v
			}
		}

		return nil
	}, nil
}
