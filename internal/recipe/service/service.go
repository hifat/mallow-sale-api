package recipeService

import (
	"context"
	"errors"
	"sync"

	inventoryHelper "github.com/hifat/mallow-sale-api/internal/inventory/helper"
	inventoryRepo "github.com/hifat/mallow-sale-api/internal/inventory/repository"
	recipeModule "github.com/hifat/mallow-sale-api/internal/recipe"
	recipeHelper "github.com/hifat/mallow-sale-api/internal/recipe/helper"
	recipeRepo "github.com/hifat/mallow-sale-api/internal/recipe/repository"
	usageUnitHelper "github.com/hifat/mallow-sale-api/internal/usageUnit/helper"
	usageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type IService interface {
	Create(ctx context.Context, req *recipeModule.Request) (*handling.ResponseItem[*recipeModule.Request], error)
	Find(ctx context.Context, query *recipeModule.QueryReq) (*handling.ResponseItems[recipeModule.Response], error)
	FindByID(ctx context.Context, id string) (*handling.ResponseItem[*recipeModule.Response], error)
	UpdateByID(ctx context.Context, id string, req *recipeModule.Request) (*handling.ResponseItem[*recipeModule.Request], error)
	DeleteByID(ctx context.Context, id string) (*handling.ResponseItem[*recipeModule.Request], error)
	UpdateNoBatch(ctx context.Context, reqs []recipeModule.UpdateOrderNoRequest) error
}

type service struct {
	logger           logger.ILogger
	recipeRepo       recipeRepo.IRepository
	inventoryRepo    inventoryRepo.IRepository
	usageUnitRepo    usageUnitRepository.IRepository
	usageUnitHelper  usageUnitHelper.IHelper
	inventoryHelper  inventoryHelper.IHelper
	recipeTypeHelper recipeHelper.IRecipeTypeHelper
}

func New(
	logger logger.ILogger,
	recipeRepo recipeRepo.IRepository,
	inventoryRepo inventoryRepo.IRepository,
	usageUnitRepo usageUnitRepository.IRepository,
	usageUnitHelper usageUnitHelper.IHelper,
	inventoryHelper inventoryHelper.IHelper,
	recipeTypeHelper recipeHelper.IRecipeTypeHelper,
) IService {
	return &service{
		logger:           logger,
		recipeRepo:       recipeRepo,
		inventoryRepo:    inventoryRepo,
		usageUnitRepo:    usageUnitRepo,
		usageUnitHelper:  usageUnitHelper,
		inventoryHelper:  inventoryHelper,
		recipeTypeHelper: recipeTypeHelper,
	}
}

func (s *service) Create(ctx context.Context, req *recipeModule.Request) (*handling.ResponseItem[*recipeModule.Request], error) {
	maxWorkers := 2
	errCh := make(chan error, maxWorkers)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		getUsageUnitName, err := s.usageUnitHelper.GetNameByCode(ctx, req.GetUsageUnitCodes())
		if err != nil {
			s.logger.Error(err)
			errCh <- err
		}

		for i, v := range req.Ingredients {
			name := getUsageUnitName(v.Unit.Code)
			if name == "" {
				errCh <- handling.ThrowErrByCode(define.CodeInvalidUsageUnit)
			}

			req.Ingredients[i].Unit.Name = name
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		getRecipeTypeByCode, err := s.recipeTypeHelper.FindAndGetByCode(ctx, []string{req.RecipeType.Code})
		if err != nil {
			s.logger.Error(err)
			errCh <- err
		}

		recipeType := getRecipeTypeByCode(req.RecipeType.Code)
		if recipeType == nil {
			errCh <- handling.ThrowErrByCode(define.CodeInvalidRecipeType)
		}

		req.RecipeType.Name = recipeType.Name
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	inventories, err := s.inventoryRepo.FindInIDs(ctx, req.GetInventoryIDs())
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	if len(inventories) != len(req.GetInventoryIDs()) {
		return nil, handling.ThrowErrByCode(define.CodeInvalidInventoryID)
	}

	err = s.recipeRepo.Create(ctx, req)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*recipeModule.Request]{
		Item: req,
	}, nil
}

func (s *service) Find(ctx context.Context, query *recipeModule.QueryReq) (*handling.ResponseItems[recipeModule.Response], error) {
	count, err := s.recipeRepo.Count(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	recipes, err := s.recipeRepo.Find(ctx, query)
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
	recipe, err := s.recipeRepo.FindByID(ctx, id)
	if err != nil {
		return nil, handling.ThrowErr(err)
	}

	getInventoryByID, err := s.inventoryHelper.FindAndGetByID(ctx, recipe.GetInventoryIDs())
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

	getRecipeTypeByCode, err := s.recipeTypeHelper.FindAndGetByCode(ctx, []string{req.RecipeType.Code})
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	recipeType := getRecipeTypeByCode(req.RecipeType.Code)
	if recipeType == nil {
		return nil, handling.ThrowErrByCode(define.CodeInvalidRecipeType)
	}

	req.RecipeType.Name = recipeType.Name

	inventories, err := s.inventoryRepo.FindInIDs(ctx, req.GetInventoryIDs())
	if err != nil {
		return nil, handling.ThrowErr(err)
	}

	if len(inventories) != len(req.GetInventoryIDs()) {
		return nil, handling.ThrowErrByCode(define.CodeInvalidInventoryID)
	}

	err = s.recipeRepo.UpdateByID(ctx, id, req)
	if err != nil {
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*recipeModule.Request]{
		Item: req,
	}, nil
}

func (s *service) UpdateNoBatch(ctx context.Context, reqs []recipeModule.UpdateOrderNoRequest) error {
	orderNoSet := make(map[int]struct{}, len(reqs))
	for _, req := range reqs {
		if _, exists := orderNoSet[req.OrderNo]; exists {
			return handling.ThrowErrByCode(define.CodeOrderNoMustBeUnique)
		}
		orderNoSet[req.OrderNo] = struct{}{}
	}

	err := s.recipeRepo.UpdateNoBatch(ctx, reqs)
	if err != nil {
		return handling.ThrowErr(err)
	}

	return nil
}

func (s *service) DeleteByID(ctx context.Context, id string) (*handling.ResponseItem[*recipeModule.Request], error) {
	_, err := s.recipeRepo.FindByID(ctx, id)
	if err != nil {
		if !errors.Is(err, define.ErrRecordNotFound) {
			s.logger.Error(err)
		}

		return nil, handling.ThrowErr(err)
	}

	err = s.recipeRepo.DeleteByID(ctx, id)
	if err != nil {
		return nil, handling.ThrowErr(err)
	}

	return nil, nil
}
