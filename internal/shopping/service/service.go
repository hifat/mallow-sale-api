package shoppingService

import (
	"context"
	"errors"
	"net/http"
	"sync"

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
	req.Status.Name = shoppingModule.EnumCodeShoppingPending.GetShoppingStatusName()

	err = s.shoppingRepo.Create(ctx, req)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*shoppingModule.Request]{
		Item: req,
	}, nil
}

func (s *service) CreateBatch(ctx context.Context, reqs []*shoppingModule.Request) (*handling.ResponseItems[*shoppingModule.Request], error) {
	if len(reqs) == 0 {
		return &handling.ResponseItems[*shoppingModule.Request]{
			Items: []*shoppingModule.Request{},
			Meta: handling.MetaResponse{
				Total: 0,
			},
		}, nil
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	numWorkers := 20
	if numWorkers > len(reqs) {
		numWorkers = len(reqs)
	}

	jobChan := make(chan *shoppingModule.Request, len(reqs))
	resChan := make(chan *shoppingModule.Request, len(reqs))
	errChan := make(chan error, len(reqs))

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()

			for req := range jobChan {
				select {
				case <-ctx.Done():
					return
				default:
				}

				res, err := s.Create(ctx, req)
				if err != nil {
					cancel()
					errChan <- err
					return
				}

				resChan <- res.Item
			}
		}()
	}

	for _, req := range reqs {
		jobChan <- req
	}
	close(jobChan)

	wg.Wait()
	close(errChan)
	close(resChan)

	if len(errChan) > 0 {
		return nil, handling.ThrowErr(<-errChan)
	}

	resShoppings := make([]*shoppingModule.Request, 0, len(resChan))
	for res := range resChan {
		resShoppings = append(resShoppings, res)
	}

	return &handling.ResponseItems[*shoppingModule.Request]{
		Items: resShoppings,
		Meta: handling.MetaResponse{
			Total: int64(len(resShoppings)),
		},
	}, nil
}

func (s *service) UpdateByID(ctx context.Context, id string, req *shoppingModule.Request) (*handling.ResponseItem[*shoppingModule.Request], error) {
	_, err := s.shoppingRepo.FindByID(ctx, id)
	if err != nil {
		if !errors.Is(define.ErrRecordNotFound, err) {
			s.logger.Error(err)
		}

		return nil, handling.ThrowErr(err)
	}

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

	err = s.shoppingRepo.UpdateByID(ctx, id, req)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*shoppingModule.Request]{
		Item: req,
	}, nil
}

func (s *service) UpdateStatus(ctx context.Context, id string, req *shoppingModule.ReqUpdateStatus) (*handling.Response, error) {
	if err := req.ValidateStatusCode(); err != nil {
		return nil, handling.ThrowErrByCode(define.CodeInvalidShoppingStatus)
	}

	_, err := s.shoppingRepo.FindByID(ctx, id)
	if err != nil {
		if !errors.Is(define.ErrRecordNotFound, err) {
			s.logger.Error(err)
		}

		return nil, handling.ThrowErr(err)
	}

	if err := s.shoppingRepo.UpdateStatus(ctx, id, req); err != nil {
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
