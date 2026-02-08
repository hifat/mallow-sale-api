package shoppingService

import (
	"context"

	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type usageUnitService struct {
	usageUnitRepository shoppingModule.IUsageUnitRepository
	logger              logger.ILogger
}

func NewUsageUnit(logger logger.ILogger, usageUnitRepository shoppingModule.IUsageUnitRepository) shoppingModule.IUsageUnitService {
	return &usageUnitService{
		usageUnitRepository,
		logger,
	}
}

func (s *usageUnitService) Create(ctx context.Context, req *shoppingModule.RequestUsageUnit) (*handling.ResponseItem[*shoppingModule.RequestUsageUnit], error) {
	err := s.usageUnitRepository.Create(ctx, req)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*shoppingModule.RequestUsageUnit]{
		Item: req,
	}, nil
}

func (s *usageUnitService) Find(ctx context.Context) (*handling.ResponseItems[shoppingModule.ResUsageUnit], error) {
	res, err := s.usageUnitRepository.Find(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItems[shoppingModule.ResUsageUnit]{
		Items: res,
	}, nil
}

func (s *usageUnitService) FindByID(ctx context.Context, id string) (*handling.ResponseItem[*shoppingModule.ResUsageUnit], error) {
	res, err := s.usageUnitRepository.FindByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*shoppingModule.ResUsageUnit]{
		Item: res,
	}, nil
}

func (s *usageUnitService) UpdateByID(ctx context.Context, id string, req *shoppingModule.RequestUsageUnit) (*handling.ResponseItem[*shoppingModule.RequestUsageUnit], error) {
	err := s.usageUnitRepository.UpdateByID(ctx, id, req)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*shoppingModule.RequestUsageUnit]{
		Item: req,
	}, nil
}

func (s *usageUnitService) DeleteByID(ctx context.Context, id string) error {
	err := s.usageUnitRepository.DeleteByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return handling.ThrowErr(err)
	}

	return nil
}
