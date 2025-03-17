package inventoryService

import (
	"context"

	"github.com/hifat/cost-calculator-api/internal/inventory"
	"github.com/hifat/cost-calculator-api/internal/inventory/inventoryRepository"
	core "github.com/hifat/goroger-core"
)

type IInventoryService interface {
	Create(ctx context.Context, req inventory.InventoryReq) error
	Find(ctx context.Context) ([]inventory.InventoryRes, error)
}

type inventoryService struct {
	inventoryRepo inventoryRepository.IInventoryRepository
	helper        core.Helper
	logger        core.Logger
}

func New(inventoryRepo inventoryRepository.IInventoryRepository, helper core.Helper, logger core.Logger) IInventoryService {
	return &inventoryService{
		inventoryRepo,
		helper,
		logger,
	}
}

func (s *inventoryService) Create(ctx context.Context, req inventory.InventoryReq) error {
	newInventory := inventory.Inventory{}
	if err := s.helper.Copy(&newInventory, req); err != nil {
		s.logger.Error(err)
		return err
	}

	if _, err := s.inventoryRepo.Create(ctx, newInventory); err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *inventoryService) Find(ctx context.Context) ([]inventory.InventoryRes, error) {
	inventories, err := s.inventoryRepo.Find(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	res := []inventory.InventoryRes{}
	if err := s.helper.Copy(&res, inventories); err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return res, nil
}
