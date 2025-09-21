package recipeHelper

import (
	"context"

	recipeModule "github.com/hifat/mallow-sale-api/internal/recipe"
	recipeRepository "github.com/hifat/mallow-sale-api/internal/recipe/repository"
)

type IHelper interface {
	FindAndGetByID(ctx context.Context, ids []string) (func(id string) *recipeModule.Response, error)
}

type helper struct {
	recipeRepository recipeRepository.IRepository
}

func New(recipeRepository recipeRepository.IRepository) IHelper {
	return &helper{
		recipeRepository: recipeRepository,
	}
}

func (h *helper) FindAndGetByID(ctx context.Context, ids []string) (func(id string) *recipeModule.Response, error) {
	recipes, err := h.recipeRepository.FindInIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	return func(id string) *recipeModule.Response {
		for _, recipe := range recipes {
			if recipe.ID == id {
				return &recipe
			}
		}

		return nil
	}, nil
}
