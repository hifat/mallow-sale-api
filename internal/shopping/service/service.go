package shoppingService

import (
	"context"
	"errors"
	"net/http"

	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
	shoppingRepository "github.com/hifat/mallow-sale-api/internal/shopping/repository"
	usageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type IService interface {
	Create(ctx context.Context, req *shoppingModule.Request) (*handling.SuccessResponse, error)
	UpdateIsComplete(ctx context.Context, req *shoppingModule.UpdateIsComplete) error
	Delete(ctx context.Context, id string) error
}

type service struct {
	logger        logger.ILogger
	shoppingRepo  shoppingRepository.IRepository
	usageUnitRepo usageUnitRepository.IRepository
}

func New(logger logger.ILogger, shoppingRepo shoppingRepository.IRepository, usageUnitRepo usageUnitRepository.IRepository) IService {
	return &service{
		logger,
		shoppingRepo,
		usageUnitRepo,
	}
}

func (s *service) Create(ctx context.Context, req *shoppingModule.Request) (*handling.SuccessResponse, error) {
	usageUnit, err := s.usageUnitRepo.FindByCode(ctx, req.PurchaseUnit.Code)
	if err != nil {
		if errors.Is(err, define.ErrRecordNotFound) {
			return nil, handling.ThrowErrByCode(define.CodeInvalidUsageUnit)
		}

		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	req.PurchaseUnit.Name = usageUnit.Name

	err = s.shoppingRepo.Create(ctx, req)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.SuccessResponse{
		Message: define.MsgCreated,
		Code:    define.CodeCreated,
		Status:  http.StatusCreated,
	}, nil
}

func (s *service) UpdateIsComplete(ctx context.Context, req *shoppingModule.UpdateIsComplete) error {
	return nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	return nil
}
