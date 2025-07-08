package service

import (
	"context"
	"errors"

	inventoryHelper "github.com/hifat/mallow-sale-api/internal/inventory/helper"
	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
	recipeModule "github.com/hifat/mallow-sale-api/internal/recipe"
	recipeRepository "github.com/hifat/mallow-sale-api/internal/recipe/repository"
	usageUnitHelper "github.com/hifat/mallow-sale-api/internal/usageUnit/helper"
	usageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type Service interface {
	Create(ctx context.Context, req *recipeModule.Request) (*handling.ResponseItem[*recipeModule.Request], error)
	Find(ctx context.Context, query *utilsModule.QueryReq) (*handling.ResponseItems[recipeModule.Response], error)
	FindByID(ctx context.Context, id string) (*handling.ResponseItem[*recipeModule.Response], error)
	UpdateByID(ctx context.Context, id string, req *recipeModule.Request) (*handling.ResponseItem[*recipeModule.Request], error)
	DeleteByID(ctx context.Context, id string) (*handling.ResponseItem[*recipeModule.Request], error)
}

type service struct {
	logger              logger.Logger
	recipeRepository    recipeRepository.Repository
	inventoryRepository inventoryRepository.Repository
	usageUnitRepository usageUnitRepository.Repository
	usageUnitHelper     usageUnitHelper.Helper
	inventoryHelper     inventoryHelper.Helper
}

func New(
	logger logger.Logger,
	recipeRepository recipeRepository.Repository,
	inventoryRepository inventoryRepository.Repository,
	usageUnitRepository usageUnitRepository.Repository,
	usageUnitHelper usageUnitHelper.Helper,
	inventoryHelper inventoryHelper.Helper,
) Service {
	return &service{
		logger:              logger,
		recipeRepository:    recipeRepository,
		inventoryRepository: inventoryRepository,
		usageUnitRepository: usageUnitRepository,
		usageUnitHelper:     usageUnitHelper,
		inventoryHelper:     inventoryHelper,
	}
}

func (s *service) Create(ctx context.Context, req *recipeModule.Request) (*handling.ResponseItem[*recipeModule.Request], error) {
	getUsageUnitName, err := s.usageUnitHelper.GetNameByCode(ctx, req.GetUsageUnitCodes())
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	for i, v := range req.Ingredients {
		name := getUsageUnitName(v.Unit.Code)
		if name == "" {
			return nil, handling.ThrowErrByCode(define.CodeInvalidUsageUnit)
		}

		req.Ingredients[i].Unit.Name = name
	}

	inventories, err := s.inventoryRepository.FindInIDs(ctx, req.GetInventoryIDs())
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	if len(inventories) != len(req.GetInventoryIDs()) {
		return nil, handling.ThrowErrByCode(define.CodeInvalidInventoryID)
	}

	err = s.recipeRepository.Create(ctx, req)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*recipeModule.Request]{
		Item: req,
	}, nil
}

func (s *service) Find(ctx context.Context, query *utilsModule.QueryReq) (*handling.ResponseItems[recipeModule.Response], error) {
	count, err := s.recipeRepository.Count(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	recipes, err := s.recipeRepository.Find(ctx, query)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItems[recipeModule.Response]{
		Items: recipes,
		Meta: handling.MetaResponse{
			Total: count,
		},
	}, nil
}

func (s *service) FindByID(ctx context.Context, id string) (*handling.ResponseItem[*recipeModule.Response], error) {
	recipe, err := s.recipeRepository.FindByID(ctx, id)
	if err != nil {
		return nil, handling.ThrowErr(err)
	}

	getInventoryByID, err := s.inventoryHelper.FindAndGetByID(ctx, recipe.GetInventoryIDFromIngredients())
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	for i, v := range recipe.Ingredients {
		inventory := getInventoryByID(v.InventoryID)
		if inventory == nil {
			s.logger.Error("inventory not found: ", v.InventoryID)
			continue
		}

		recipe.Ingredients[i].Inventory = &inventory.Prototype
	}

	return &handling.ResponseItem[*recipeModule.Response]{
		Item: recipe,
	}, nil
}

func (s *service) UpdateByID(ctx context.Context, id string, req *recipeModule.Request) (*handling.ResponseItem[*recipeModule.Request], error) {
	getUsageUnitName, err := s.usageUnitHelper.GetNameByCode(ctx, req.GetUsageUnitCodes())
	if err != nil {
		return nil, handling.ThrowErr(err)
	}

	for i, v := range req.Ingredients {
		name := getUsageUnitName(v.Unit.Code)
		if name == "" {
			return nil, handling.ThrowErrByCode(define.CodeInvalidUsageUnit)
		}

		req.Ingredients[i].Unit.Name = name
	}

	inventories, err := s.inventoryRepository.FindInIDs(ctx, req.GetInventoryIDs())
	if err != nil {
		return nil, handling.ThrowErr(err)
	}

	if len(inventories) != len(req.GetInventoryIDs()) {
		return nil, handling.ThrowErrByCode(define.CodeInvalidInventoryID)
	}

	err = s.recipeRepository.UpdateByID(ctx, id, req)
	if err != nil {
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*recipeModule.Request]{
		Item: req,
	}, nil
}

func (s *service) DeleteByID(ctx context.Context, id string) (*handling.ResponseItem[*recipeModule.Request], error) {
	_, err := s.recipeRepository.FindByID(ctx, id)
	if err != nil {
		if !errors.Is(err, define.ErrRecordNotFound) {
			s.logger.Error(err)
		}

		return nil, handling.ThrowErr(err)
	}

	err = s.recipeRepository.DeleteByID(ctx, id)
	if err != nil {
		return nil, handling.ThrowErr(err)
	}

	return nil, nil
}
