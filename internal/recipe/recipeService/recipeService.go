package recipeService

import (
	"context"

	"github.com/hifat/cost-calculator-api/internal/recipe"
	"github.com/hifat/cost-calculator-api/internal/recipe/recipeRepository"
	core "github.com/hifat/goroger-core"
	"github.com/hifat/goroger-core/rules"
)

type IRecipeService interface {
	Create(ctx context.Context, req recipe.RecipeReq) (*recipe.RecipeRes, error)
	Find(ctx context.Context) ([]recipe.RecipeRes, error)
	FindByID(ctx context.Context, id string) (*recipe.RecipeRes, error)
	Update(ctx context.Context, id string, req recipe.RecipeReq) error
	Delete(ctx context.Context, id string) error
}

type recipeService struct {
	logger     core.Logger
	validator  rules.Validator
	recipeRepo recipeRepository.IRecipeRepository
}

func New(logger core.Logger, validator rules.Validator, recipeRepo recipeRepository.IRecipeRepository) IRecipeService {
	return &recipeService{
		logger,
		validator,
		recipeRepo,
	}
}

func (s *recipeService) Create(ctx context.Context, req recipe.RecipeReq) (*recipe.RecipeRes, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, err
	}

	recipeID, err := s.recipeRepo.Create(ctx, req)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	res, err := s.recipeRepo.FindByID(ctx, recipeID)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return res, nil
}

func (s *recipeService) Find(ctx context.Context) ([]recipe.RecipeRes, error) {
	res := []recipe.RecipeRes{}
	res, err := s.recipeRepo.Find(ctx)
	if err != nil {
		s.logger.Error(err)
		return res, err
	}

	return res, nil
}

func (s *recipeService) FindByID(ctx context.Context, id string) (*recipe.RecipeRes, error) {
	res, err := s.recipeRepo.FindByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return res, nil
}

func (s *recipeService) Update(ctx context.Context, id string, req recipe.RecipeReq) error {
	if err := s.recipeRepo.Update(ctx, id, req); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *recipeService) Delete(ctx context.Context, id string) error {
	if err := s.recipeRepo.Delete(ctx, id); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}
