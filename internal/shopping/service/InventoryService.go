package shoppingService

import (
	"context"
	"errors"

	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
	supplierRepository "github.com/hifat/mallow-sale-api/internal/supplier/repository"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type inventoryService struct {
	repo          shoppingModule.IInventoryRepository
	inventoryRepo inventoryRepository.IRepository
	supplierRepo  supplierRepository.IRepository
	logger        logger.ILogger
}

func NewInventory(repo shoppingModule.IInventoryRepository, inventoryRepo inventoryRepository.IRepository, supplierRepo supplierRepository.IRepository, logger logger.ILogger) shoppingModule.IInventoryService {
	return &inventoryService{
		repo:          repo,
		inventoryRepo: inventoryRepo,
		supplierRepo:  supplierRepo,
		logger:        logger,
	}
}

func (s *inventoryService) Create(ctx context.Context, req *shoppingModule.RequestShoppingInventory) (*handling.ResponseItem[*shoppingModule.RequestShoppingInventory], error) {
	inventory, err := s.inventoryRepo.FindByID(ctx, req.InventoryID)
	if err != nil {
		if errors.Is(err, define.ErrRecordNotFound) {
			return nil, handling.ThrowErrByCode(define.CodeInvalidInventoryID)
		}

		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	req.InventoryName = inventory.Name

	supplier, err := s.supplierRepo.FindByID(ctx, req.SupplierID)
	if err != nil {
		if errors.Is(err, define.ErrRecordNotFound) {
			return nil, handling.ThrowErrByCode(define.CodeInvalidSupplierID)
		}

		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	req.SupplierName = supplier.Name

	if err := s.repo.Create(ctx, req); err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*shoppingModule.RequestShoppingInventory]{
		Item: req,
	}, nil
}

func (s *inventoryService) Find(ctx context.Context) (*handling.ResponseItems[shoppingModule.ResShoppingInventory], error) {
	shoppingInvs, err := s.repo.Find(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItems[shoppingModule.ResShoppingInventory]{
		Items: shoppingInvs,
		Meta: handling.MetaResponse{
			Total: int64(len(shoppingInvs)),
		},
	}, nil
}

func (s *inventoryService) DeleteByID(ctx context.Context, id string) error {
	if err := s.repo.DeleteByID(ctx, id); err != nil {
		s.logger.Error(err)
		return handling.ThrowErr(err)
	}

	return nil
}
