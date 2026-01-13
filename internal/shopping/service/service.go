package shoppingService

import (
	"context"
	"errors"
	"net/http"

	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
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
}

func New(logger logger.ILogger, shoppingRepo shoppingModule.IRepository, usageUnitRepo usageUnitRepository.IRepository, usageUnitHelper usageUnitHelper.IHelper) shoppingModule.IService {
	return &service{
		logger,
		shoppingRepo,
		usageUnitRepo,
		usageUnitHelper,
	}
}

func (s *service) Find(ctx context.Context) (*handling.ResponseItems[shoppingModule.Response], error) {
	res := []shoppingModule.Response{}
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

func (s *service) Create(ctx context.Context, req *shoppingModule.Request) (*handling.ResponseItem[*shoppingModule.Request], error) {
	getNameByCode, err := s.usageUnitHelper.GetNameByCode(ctx, req.GetPurchaseUnitCodes())
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	for i, v := range req.Inventories {
		req.Inventories[i].PurchaseUnit.Name = getNameByCode(v.PurchaseUnit.Code)
	}

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
