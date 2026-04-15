package inventoryService

import (
	"context"
	"errors"
	"sync"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	pricePresetModule "github.com/hifat/mallow-sale-api/internal/pricePreset"
	recipeModule "github.com/hifat/mallow-sale-api/internal/recipe"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"github.com/hifat/mallow-sale-api/pkg/utils"
)

type IService interface {
	Create(ctx context.Context, req *inventoryModule.Request) (*handling.ResponseItem[*inventoryModule.Request], error)
	Find(ctx context.Context, query *utilsModule.QueryReq) (*handling.ResponseItems[inventoryModule.Response], error)
	FindByID(ctx context.Context, id string) (*handling.ResponseItem[*inventoryModule.Response], error)
	UpdateByID(ctx context.Context, id string, req *inventoryModule.Request) (*handling.ResponseItem[*inventoryModule.Request], error)
	UpdatePurchasePriceByPreset(ctx context.Context, id string, req *inventoryModule.UpdatePresetPriceReq) (*handling.ResponseItem[*inventoryModule.Response], error)
	DeleteByID(ctx context.Context, id string) error
}

type service struct {
	mu              sync.Mutex
	logger          logger.ILogger
	inventoryRepo   inventoryModule.IRepository
	usageUnitRepo   usageUnitModule.IRepository
	pricePresetRepo pricePresetModule.IRepository
	recipeRepo      recipeModule.IRepository
}

func New(
	logger logger.ILogger,
	inventoryRepo inventoryModule.IRepository,
	usageUnitRepo usageUnitModule.IRepository,
	pricePresetRepo pricePresetModule.IRepository,
	recipeRepo recipeModule.IRepository,
) IService {
	return &service{
		logger:          logger,
		inventoryRepo:   inventoryRepo,
		usageUnitRepo:   usageUnitRepo,
		pricePresetRepo: pricePresetRepo,
		recipeRepo:      recipeRepo,
	}
}

func (s *service) Create(ctx context.Context, req *inventoryModule.Request) (*handling.ResponseItem[*inventoryModule.Request], error) {
	numWorkers := 2
	errCh := make(chan error, numWorkers)

	var wg sync.WaitGroup
	wg.Add(numWorkers)
	go func() {
		defer wg.Done()
		inventory, err := s.inventoryRepo.FindByName(ctx, req.Name)
		if err != nil {
			if !errors.Is(err, define.ErrRecordNotFound) {
				s.logger.Error(err)
				errCh <- handling.ThrowErr(err)
				return
			}
		}

		if inventory != nil {
			errCh <- handling.ThrowErrByCode(define.CodeDuplicatedInventoryName)
		}
	}()

	go func() {
		defer wg.Done()
		usageUnit, err := s.usageUnitRepo.FindByCode(ctx, req.PurchaseUnit.Code)
		if err != nil {
			if !errors.Is(err, define.ErrRecordNotFound) {
				s.logger.Error(err)
				errCh <- handling.ThrowErr(err)
			}

			errCh <- handling.ThrowErrByCode(define.CodeInvalidUsageUnit)
			return
		}

		s.mu.Lock()
		req.PurchaseUnit.Name = usageUnit.Name
		s.mu.Unlock()
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

	err := s.inventoryRepo.Create(ctx, req)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*inventoryModule.Request]{
		Item: req,
	}, nil
}

func (s *service) Find(ctx context.Context, query *utilsModule.QueryReq) (*handling.ResponseItems[inventoryModule.Response], error) {
	count, err := s.inventoryRepo.Count(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	inventories, err := s.inventoryRepo.Find(ctx, query)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItems[inventoryModule.Response]{
		Items: inventories,
		Meta: handling.MetaResponse{
			Total: count,
		},
	}, nil
}

func (s *service) FindByID(ctx context.Context, id string) (*handling.ResponseItem[*inventoryModule.Response], error) {
	inventory, err := s.inventoryRepo.FindByID(ctx, id)
	if err != nil {
		if !errors.Is(err, define.ErrRecordNotFound) {
			s.logger.Error(err)
		}

		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*inventoryModule.Response]{
		Item: inventory,
	}, nil
}

func (s *service) UpdateByID(ctx context.Context, id string, req *inventoryModule.Request) (*handling.ResponseItem[*inventoryModule.Request], error) {
	_, err := s.inventoryRepo.FindByID(ctx, id)
	if err != nil {
		if !errors.Is(err, define.ErrRecordNotFound) {
			s.logger.Error(err)
		}

		return nil, handling.ThrowErr(err)
	}

	usageUnit, err := s.usageUnitRepo.FindByCode(ctx, req.PurchaseUnit.Code)
	if err != nil {
		if !errors.Is(err, define.ErrRecordNotFound) {
			s.logger.Error(err)
			return nil, handling.ThrowErr(err)
		}
	}

	if usageUnit == nil {
		return nil, handling.ThrowErrByCode(define.CodeInvalidUsageUnit)
	}

	req.PurchaseUnit.Name = usageUnit.Name

	err = s.inventoryRepo.UpdateByID(ctx, id, req)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*inventoryModule.Request]{
		Item: req,
	}, nil
}

func (s *service) UpdatePurchasePriceByPreset(ctx context.Context, id string, req *inventoryModule.UpdatePresetPriceReq) (*handling.ResponseItem[*inventoryModule.Response], error) {
	preset, err := s.pricePresetRepo.FindByPriceID(ctx, req.PresetPriceID)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	var priceValue float64
	var found bool
	for _, p := range preset.Prices {
		if p.ID == req.PresetPriceID {
			priceValue = p.Price
			found = true
			break
		}
	}

	if !found {
		return nil, handling.ThrowErrByCode(define.CodeInvalidPricePresetID)
	}

	if preset.InventoryID != id {
		return nil, handling.ThrowErrByCode(define.CodeInvalidInventoryID)
	}

	_, err = s.inventoryRepo.FindByID(ctx, id)
	if err != nil {
		if !errors.Is(err, define.ErrRecordNotFound) {
			s.logger.Error(err)
		}
		return nil, handling.ThrowErr(err)
	}

	err = s.inventoryRepo.UpdatePurchasePrice(ctx, id, priceValue)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	if syncErr := s.syncRecipeCosts(ctx, id); syncErr != nil {
		s.logger.Error(syncErr)
	}

	updatedInventory, err := s.inventoryRepo.FindByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*inventoryModule.Response]{
		Item: updatedInventory,
	}, nil
}

// syncRecipeCosts finds all recipes that use inventoryID and recalculates
// each recipe's cost using the same logic as recipe.Create (// Total Cost).
func (s *service) syncRecipeCosts(ctx context.Context, inventoryID string) error {
	recipes, err := s.recipeRepo.FindByInventoryID(ctx, inventoryID)
	if err != nil {
		return err
	}

	for _, recipe := range recipes {
		inventoryIDs := recipe.GetInventoryIDs()

		inventories, err := s.inventoryRepo.FindInIDs(ctx, inventoryIDs)
		if err != nil {
			return err
		}

		inventoryByID := make(map[string]*inventoryModule.Response, len(inventories))
		for i := range inventories {
			inventoryByID[inventories[i].ID] = &inventories[i]
		}

		// Total Cost — same logic as recipe/service/service.go Create
		var newCost float64
		for _, ingredient := range recipe.Ingredients {
			inventory := inventoryByID[ingredient.InventoryID]
			if inventory == nil {
				continue
			}

			actualPrice := utils.CalculateActualPrice(inventory.PurchasePrice, float64(inventory.YieldPercentage))
			pricePerUnit := actualPrice / inventory.PurchaseQuantity
			newCost += pricePerUnit * float64(ingredient.Quantity)
		}

		if err := s.recipeRepo.UpdateCost(ctx, recipe.ID, newCost); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) DeleteByID(ctx context.Context, id string) error {
	_, err := s.inventoryRepo.FindByID(ctx, id)
	if err != nil {
		if !errors.Is(err, define.ErrRecordNotFound) {
			s.logger.Error(err)
		}

		return handling.ThrowErr(err)
	}

	err = s.inventoryRepo.DeleteByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return handling.ThrowErr(err)
	}

	return nil
}
