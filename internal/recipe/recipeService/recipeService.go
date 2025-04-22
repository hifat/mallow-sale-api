package recipeService

import (
	"context"
	"errors"
	"fmt"

	"github.com/hifat/cost-calculator-api/internal/inventory"
	"github.com/hifat/cost-calculator-api/internal/inventory/inventoryRepository"
	"github.com/hifat/cost-calculator-api/internal/recipe"
	"github.com/hifat/cost-calculator-api/internal/recipe/recipeRepository"
	usageUnitServiceUtils "github.com/hifat/cost-calculator-api/pkg/utils/serviceUtils"
	core "github.com/hifat/goroger-core"
	"github.com/hifat/goroger-core/rules"
)

type IRecipeService interface {
	Create(ctx context.Context, req recipe.RecipeReq) (*recipe.RecipeRes, error)
	Find(ctx context.Context) ([]recipe.RecipeRes, error)
	FindByID(ctx context.Context, id string) (*recipe.RecipeRes, error)
	Update(ctx context.Context, id string, req recipe.UpdateRecipeReq) error
	Delete(ctx context.Context, id string) error
}

type recipeService struct {
	logger           core.Logger
	validator        rules.Validator
	helper           core.Helper
	usageServiceUtil usageUnitServiceUtils.IUsageUnitServiceUtils
	recipeRepo       recipeRepository.IRecipeRepository
	inventoryRepo    inventoryRepository.IInventoryRepository
}

func New(logger core.Logger, validator rules.Validator, helper core.Helper, usageServiceUtil usageUnitServiceUtils.IUsageUnitServiceUtils, recipeRepo recipeRepository.IRecipeRepository, inventoryRepo inventoryRepository.IInventoryRepository) IRecipeService {
	return &recipeService{
		logger,
		validator,
		helper,
		usageServiceUtil,
		recipeRepo,
		inventoryRepo,
	}
}

func (s *recipeService) Create(ctx context.Context, req recipe.RecipeReq) (*recipe.RecipeRes, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, err
	}

	usageUnitCodes := make([]string, 0, len(req.Inventories))
	for _, inventory := range req.Inventories {
		usageUnitCodes = append(usageUnitCodes, inventory.UsageUnitCode)
	}

	mapUsageUnit, err := s.usageServiceUtil.MapUsageUnitName(ctx, usageUnitCodes)
	if err != nil {
		return nil, err
	}

	for i, inventory := range req.Inventories {
		usageUnitName, ok := mapUsageUnit[inventory.UsageUnitCode]
		if !ok {
			return nil, errors.New("invalid usageUnitName")
		}

		req.Inventories[i].UsageUnit.SetAttr(inventory.UsageUnitCode, usageUnitName)
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

	inventoryIDs := make([]string, 0, len(res.Inventories))
	for _, inventory := range res.Inventories {
		inventoryIDs = append(inventoryIDs, inventory.InventoryID)
	}

	inventories, err := s.inventoryRepo.FindInID(ctx, inventoryIDs)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	mapInventory := map[string]inventory.Inventory{}
	for _, _inventory := range inventories {
		mapInventory[_inventory.ID] = _inventory
	}

	for i, _inventory := range res.Inventories {
		inv, ok := mapInventory[_inventory.InventoryID]
		if !ok {
			s.logger.Warn(fmt.Sprintf("not found inventory id: %s", _inventory.InventoryID))
			continue
		}

		res.Inventories[i].Inventory = &inventory.InventoryPrototype{}
		if err := s.helper.Copy(&res.Inventories[i].Inventory, inv); err != nil {
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

	usageUnitCodes := make([]string, 0, len(req.Inventories))
	for _, inventory := range req.Inventories {
		usageUnitCodes = append(usageUnitCodes, inventory.UsageUnitCode)
	}

	mapUsageUnit, err := s.usageServiceUtil.MapUsageUnitName(ctx, usageUnitCodes)
	if err != nil {
		return err
	}

	for i, inventory := range req.Inventories {
		usageUnitName, ok := mapUsageUnit[inventory.UsageUnitCode]
		if !ok {
			return errors.New("invalid usageUnitName")
		}

		req.Inventories[i].UsageUnit.SetAttr(inventory.UsageUnitCode, usageUnitName)
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
