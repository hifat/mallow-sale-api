package recipeService

import (
	"context"
	"errors"
	"fmt"

	core "github.com/hifat/goroger-core"
	"github.com/hifat/goroger-core/rules"
	"github.com/hifat/mallow-sale-api/internal/inventory"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryRepository"
	"github.com/hifat/mallow-sale-api/internal/recipe"
	"github.com/hifat/mallow-sale-api/internal/recipe/recipeRepository"
	usageUnitServiceUtils "github.com/hifat/mallow-sale-api/pkg/utils/serviceUtils"
)

type IRecipeService interface {
	Create(ctx context.Context, req recipe.RecipeReq) (*recipe.RecipeRes, error)
	Find(ctx context.Context) ([]recipe.RecipeRes, error)
	FindByID(ctx context.Context, id string) (*recipe.RecipeRes, error)
	Update(ctx context.Context, id string, req recipe.UpdateRecipeReq) error
	Delete(ctx context.Context, id string) error
}

type recipeService struct {
	logger            core.Logger
	validator         rules.Validator
	helper            core.Helper
	usageServiceUtil  usageUnitServiceUtils.IUsageUnitServiceUtils
	recipeRepo        recipeRepository.IRecipeRepository
	inventoryRepo     inventoryRepository.IInventoryRepository
	inventoryGRPCRepo inventoryRepository.IInventoryGRPCRepository
}

func New(
	logger core.Logger,
	validator rules.Validator,
	helper core.Helper,
	usageServiceUtil usageUnitServiceUtils.IUsageUnitServiceUtils,
	recipeRepo recipeRepository.IRecipeRepository,
	inventoryRepo inventoryRepository.IInventoryRepository,
	inventoryGRPCRepo inventoryRepository.IInventoryGRPCRepository,
) IRecipeService {
	return &recipeService{
		logger,
		validator,
		helper,
		usageServiceUtil,
		recipeRepo,
		inventoryRepo,
		inventoryGRPCRepo,
	}
}

func (s *recipeService) Create(ctx context.Context, req recipe.RecipeReq) (*recipe.RecipeRes, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, err
	}

	// TODO: Enable this code when finish migration
	// usageUnitCodes := make([]string, 0, len(req.Ingredients))
	// for _, ingredient := range req.Ingredients {
	// 	usageUnitCodes = append(usageUnitCodes, ingredient.UsageUnitCode)
	// }

	// mapUsageUnit, err := s.usageServiceUtil.MapUsageUnitName(ctx, usageUnitCodes)
	// if err != nil {
	// 	return nil, err
	// }

	// for i, ingredient := range req.Ingredients {
	// 	usageUnitName, ok := mapUsageUnit[ingredient.UsageUnitCode]
	// 	if !ok {
	// 		return nil, errors.New("invalid usageUnitName")
	// 	}

	// 	req.Ingredients[i].UsageUnit.SetAttr(ingredient.UsageUnitCode, usageUnitName)
	// }

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

	inventoryIDs := make([]string, 0, len(res.Ingredients))
	for _, inventory := range res.Ingredients {
		inventoryIDs = append(inventoryIDs, inventory.InventoryID)
	}

	inventories, err := s.inventoryGRPCRepo.FindInID(ctx, inventoryIDs)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	mapInventory := map[string]inventory.Inventory{}
	for _, _inventory := range inventories {
		mapInventory[_inventory.ID] = _inventory
	}

	for i, _inventory := range res.Ingredients {
		inv, ok := mapInventory[_inventory.InventoryID]
		if !ok {
			s.logger.Warn(fmt.Sprintf("not found inventory id: %s", _inventory.InventoryID))
			continue
		}

		res.Ingredients[i].Inventory = &inventory.InventoryPrototype{}
		if err := s.helper.Copy(&res.Ingredients[i].Inventory, inv); err != nil {
			s.logger.Error(err)
			return nil, err
		}
	}

	return res, nil
}

func (s *recipeService) Update(ctx context.Context, id string, req recipe.UpdateRecipeReq) error {
	if err := s.validator.Validate(req); err != nil {
		return err
	}

	usageUnitCodes := make([]string, 0, len(req.Ingredients))
	for _, inventory := range req.Ingredients {
		usageUnitCodes = append(usageUnitCodes, inventory.UsageUnitCode)
	}

	mapUsageUnit, err := s.usageServiceUtil.MapUsageUnitName(ctx, usageUnitCodes)
	if err != nil {
		return err
	}

	for i, inventory := range req.Ingredients {
		usageUnitName, ok := mapUsageUnit[inventory.UsageUnitCode]
		if !ok {
			return errors.New("invalid usageUnitName")
		}

		req.Ingredients[i].UsageUnit.SetAttr(inventory.UsageUnitCode, usageUnitName)
	}

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
