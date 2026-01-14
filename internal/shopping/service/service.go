package shoppingService

import (
	"context"
	"errors"
	"net/http"

	inventoryHelper "github.com/hifat/mallow-sale-api/internal/inventory/helper"
	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
	supplierHelper "github.com/hifat/mallow-sale-api/internal/supplier/helper"
	usageUnitHelper "github.com/hifat/mallow-sale-api/internal/usageUnit/helper"
	usageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type service struct {
	logger          logger.ILogger
	shoppingRepo    shoppingModule.IRepository
	usageUnitRepo   usageUnitRepository.IRepository
	usageUnitHelper usageUnitHelper.IHelper
	supplierHelper  supplierHelper.IHelper
	inventoryHelper inventoryHelper.IHelper
}

func New(
	logger logger.ILogger,
	shoppingRepo shoppingModule.IRepository,
	usageUnitRepo usageUnitRepository.IRepository,
	usageUnitHelper usageUnitHelper.IHelper,
	supplierHelper supplierHelper.IHelper,
	inventoryHelper inventoryHelper.IHelper,
) shoppingModule.IService {
	return &service{
		logger,
		shoppingRepo,
		usageUnitRepo,
		usageUnitHelper,
		supplierHelper,
		inventoryHelper,
	}
}

func (s *service) Find(ctx context.Context) (*handling.ResponseItems[shoppingModule.Response], error) {
	res, err := s.shoppingRepo.Find(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItems[shoppingModule.Response]{
		Items: res,
		Meta: handling.MetaResponse{
			Total: int64(len(res)),
		},
	}, nil
}

func (s *service) FindByID(ctx context.Context, id string) (*handling.ResponseItem[*shoppingModule.Response], error) {
	res, err := s.shoppingRepo.FindByID(ctx, id)
	if err != nil {
		if !errors.Is(define.ErrRecordNotFound, err) {
			s.logger.Error(err)
		}

		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*shoppingModule.Response]{
		Item: res,
	}, nil
}

func (s *service) Create(ctx context.Context, req *shoppingModule.Request) (*handling.ResponseItem[*shoppingModule.Request], error) {
	findSupplierByID, err := s.supplierHelper.FindAndGetByID(ctx, []string{req.SupplierID})
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	supplier := findSupplierByID(req.SupplierID)
	if supplier == nil {
		return nil, handling.ThrowErrByCode(define.CodeInvalidSupplierID)
	}

	req.SupplierName = supplier.Name

	getNameByCode, err := s.usageUnitHelper.GetNameByCode(ctx, req.GetPurchaseUnitCodes())
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	findInventoryByID, err := s.inventoryHelper.FindAndGetByID(ctx, req.GetInventoryIDs())
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	for i, v := range req.Inventories {
		inventory := findInventoryByID(v.InventoryID)
		if inventory == nil {
			return nil, handling.ThrowErrByCode(define.CodeInvalidInventoryID)
		}

		req.Inventories[i].InventoryName = inventory.Name

		purchaseUnitName := getNameByCode(v.PurchaseUnit.Code)
		if purchaseUnitName == "" {
			return nil, handling.ThrowErrByCode(define.CodeInvalidPurchaseUnit)
		}

		req.Inventories[i].PurchaseUnit.Name = purchaseUnitName

		req.Inventories[i].Status = shoppingModule.InventoryStatus{
			Code: shoppingModule.EnumCodeInventoryPending,
			Name: v.Status.Name,
		}
	}

	req.Status.Code = shoppingModule.EnumCodeShoppingPending

	err = s.shoppingRepo.Create(ctx, req)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*shoppingModule.Request]{
		Item: req,
	}, nil
}

// TODO: Wait destroy this function
func (s *service) UpdateIsComplete(ctx context.Context, id string, req *shoppingModule.ReqUpdateIsComplete) (*handling.Response, error) {
	_, err := s.shoppingRepo.FindByID(ctx, id)
	if err != nil {
		if !errors.Is(define.ErrRecordNotFound, err) {
			s.logger.Error(err)
		}

		return nil, handling.ThrowErr(err)
	}

	if err := s.shoppingRepo.UpdateIsComplete(ctx, id, req); err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.Response{
		Message: define.MsgUpdated,
		Code:    define.CodeUpdated,
		Status:  http.StatusOK,
	}, nil
}

func (s *service) DeleteByID(ctx context.Context, id string) (*handling.Response, error) {
	_, err := s.shoppingRepo.FindByID(ctx, id)
	if err != nil {
		if !errors.Is(define.ErrRecordNotFound, err) {
			s.logger.Error(err)
		}

		return nil, handling.ThrowErr(err)
	}

	if err := s.shoppingRepo.DeleteByID(ctx, id); err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.Response{
		Message: define.MsgDeleted,
		Code:    define.CodeDeleted,
		Status:  http.StatusOK,
	}, nil
}
