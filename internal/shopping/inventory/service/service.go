package shoppingInventoryModule

import (
	"context"
	"errors"

	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
	shoppingInventoryModule "github.com/hifat/mallow-sale-api/internal/shopping/inventory"
	supplierRepository "github.com/hifat/mallow-sale-api/internal/supplier/repository"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type service struct {
	repo          shoppingInventoryModule.IRepository
	inventoryRepo inventoryRepository.IRepository
	supplierRepo  supplierRepository.IRepository
	logger        logger.ILogger
}

func New(repo shoppingInventoryModule.IRepository, inventoryRepo inventoryRepository.IRepository, supplierRepo supplierRepository.IRepository, logger logger.ILogger) shoppingInventoryModule.IService {
	return &service{
		repo:          repo,
		inventoryRepo: inventoryRepo,
		supplierRepo:  supplierRepo,
		logger:        logger,
	}
}

func (s *service) Create(ctx context.Context, req *shoppingInventoryModule.Request) (*handling.ResponseItem[*shoppingInventoryModule.Request], error) {
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
			return nil, handling.ThrowErrByCode(define.CodeInvalidInventoryID)
		}

		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	req.SupplierName = supplier.Name

	if err := s.repo.Create(ctx, req); err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*shoppingInventoryModule.Request]{
		Item: req,
	}, nil
}

func (s *service) Find(ctx context.Context) (*handling.ResponseItems[shoppingInventoryModule.Response], error) {
	invShoppings, err := s.repo.Find(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItems[shoppingInventoryModule.Response]{
		Items: invShoppings,
	}, nil
}

func (s *service) DeleteByID(ctx context.Context, id string) error {
	if err := s.repo.DeleteByID(ctx, id); err != nil {
		s.logger.Error(err)
		return handling.ThrowErr(err)
	}

	return nil
}
