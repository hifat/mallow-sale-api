package stockService

import (
	"context"
	"errors"

	inventoryHelper "github.com/hifat/mallow-sale-api/internal/inventory/helper"
	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
	stockModule "github.com/hifat/mallow-sale-api/internal/stock"
	stockRepository "github.com/hifat/mallow-sale-api/internal/stock/repository"
	supplierHelper "github.com/hifat/mallow-sale-api/internal/supplier/helper"
	supplierRepository "github.com/hifat/mallow-sale-api/internal/supplier/repository"
	usageUnitHelper "github.com/hifat/mallow-sale-api/internal/usageUnit/helper"
	usageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type Service interface {
	Create(ctx context.Context, req *stockModule.Request) (*handling.ResponseItem[*stockModule.Request], error)
	Find(ctx context.Context, query *utilsModule.QueryReq) (*handling.ResponseItems[stockModule.Response], error)
	FindByID(ctx context.Context, id string) (*handling.ResponseItem[*stockModule.Response], error)
	UpdateByID(ctx context.Context, id string, req *stockModule.Request) (*handling.ResponseItem[*stockModule.Request], error)
	DeleteByID(ctx context.Context, id string) error
}

type service struct {
	stockRepository     stockRepository.Repository
	inventoryRepository inventoryRepository.Repository
	supplierRepository  supplierRepository.Repository
	usageUnitRepository usageUnitRepository.Repository
	inventoryHelper     inventoryHelper.Helper
	supplierHelper      supplierHelper.Helper
	usageUnitHelper     usageUnitHelper.Helper
	logger              logger.Logger
}

func New(
	stockRepository stockRepository.Repository,
	inventoryRepository inventoryRepository.Repository,
	supplierRepository supplierRepository.Repository,
	usageUnitRepository usageUnitRepository.Repository,
	inventoryHelper inventoryHelper.Helper,
	supplierHelper supplierHelper.Helper,
	usageUnitHelper usageUnitHelper.Helper,
	logger logger.Logger,
) Service {
	return &service{
		stockRepository:     stockRepository,
		inventoryRepository: inventoryRepository,
		supplierRepository:  supplierRepository,
		usageUnitRepository: usageUnitRepository,
		inventoryHelper:     inventoryHelper,
		supplierHelper:      supplierHelper,
		usageUnitHelper:     usageUnitHelper,
		logger:              logger,
	}
}

func (s *service) validateStockRequest(ctx context.Context, req *stockModule.Request) error {
	// Validate usage unit and get name
	getUsageUnitName, err := s.usageUnitHelper.GetNameByCode(ctx, []string{req.PurchaseUnit.Code})
	if err != nil {
		s.logger.Error(err)
		return handling.ThrowErr(err)
	}

	usageUnitName := getUsageUnitName(req.PurchaseUnit.Code)
	if usageUnitName == "" {
		return handling.ThrowErrByCode(define.CodeInvalidUsageUnit)
	}
	req.PurchaseUnit.Name = usageUnitName

	// Validate inventory ID using helper
	getInventoryByID, err := s.inventoryHelper.FindAndGetByID(ctx, []string{req.InventoryID})
	if err != nil {
		s.logger.Error(err)
		return handling.ThrowErr(err)
	}

	inventory := getInventoryByID(req.InventoryID)
	if inventory == nil {
		return handling.ThrowErrByCode(define.CodeInvalidInventoryID)
	}

	// Validate supplier ID using helper
	getSupplierByID, err := s.supplierHelper.FindAndGetByID(ctx, []string{req.SupplierID})
	if err != nil {
		s.logger.Error(err)
		return handling.ThrowErr(err)
	}

	supplier := getSupplierByID(req.SupplierID)
	if supplier == nil {
		return handling.ThrowErrByCode(define.CodeInvalidSupplierID)
	}

	return nil
}

func (s *service) Create(ctx context.Context, req *stockModule.Request) (*handling.ResponseItem[*stockModule.Request], error) {
	if err := s.validateStockRequest(ctx, req); err != nil {
		return nil, err
	}

	err := s.stockRepository.Create(ctx, req)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}
	return &handling.ResponseItem[*stockModule.Request]{Item: req}, nil
}

func (s *service) Find(ctx context.Context, query *utilsModule.QueryReq) (*handling.ResponseItems[stockModule.Response], error) {
	count, err := s.stockRepository.Count(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	stocks, err := s.stockRepository.Find(ctx, query)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItems[stockModule.Response]{
		Items: stocks,
		Meta:  handling.MetaResponse{Total: count},
	}, nil
}

func (s *service) FindByID(ctx context.Context, id string) (*handling.ResponseItem[*stockModule.Response], error) {
	stock, err := s.stockRepository.FindByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	// Populate Inventory
	if stock.InventoryID != "" {
		getInventoryByID, err := s.inventoryHelper.FindAndGetByID(ctx, []string{stock.InventoryID})
		if err != nil {
			s.logger.Error(err)
			return nil, handling.ThrowErr(err)
		}
		inventory := getInventoryByID(stock.InventoryID)
		if inventory != nil {
			stock.Inventory = &inventory.Prototype
		}
	}

	// Populate Supplier
	if stock.SupplierID != "" {
		getSupplierByID, err := s.supplierHelper.FindAndGetByID(ctx, []string{stock.SupplierID})
		if err != nil {
			s.logger.Error(err)
			return nil, handling.ThrowErr(err)
		}
		supplier := getSupplierByID(stock.SupplierID)
		if supplier != nil {
			stock.Supplier = &supplier.Prototype
		}
	}

	return &handling.ResponseItem[*stockModule.Response]{Item: stock}, nil
}

func (s *service) UpdateByID(ctx context.Context, id string, req *stockModule.Request) (*handling.ResponseItem[*stockModule.Request], error) {
	// Check if stock exists
	_, err := s.stockRepository.FindByID(ctx, id)
	if err != nil {
		if !errors.Is(err, define.ErrRecordNotFound) {
			s.logger.Error(err)
		}
		return nil, handling.ThrowErr(err)
	}

	if err := s.validateStockRequest(ctx, req); err != nil {
		return nil, err
	}

	err = s.stockRepository.UpdateByID(ctx, id, req)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*stockModule.Request]{Item: req}, nil
}

func (s *service) DeleteByID(ctx context.Context, id string) error {
	_, err := s.stockRepository.FindByID(ctx, id)
	if err != nil {
		if !errors.Is(err, define.ErrRecordNotFound) {
			s.logger.Error(err)
		}
		return handling.ThrowErr(err)
	}

	err = s.stockRepository.DeleteByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return handling.ThrowErr(err)
	}

	return nil
}
