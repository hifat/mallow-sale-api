package inventoryService

import (
	"context"
	"errors"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
	usageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type Service interface {
	Create(ctx context.Context, req *inventoryModule.Request) (*handling.ResponseItem[*inventoryModule.Request], error)
	FindByID(ctx context.Context, id string) (*handling.ResponseItem[*inventoryModule.Response], error)
	Find(ctx context.Context) (*handling.ResponseItems[inventoryModule.Response], error)
	UpdateByID(ctx context.Context, id string, req *inventoryModule.Request) (*handling.ResponseItem[*inventoryModule.Request], error)
	DeleteByID(ctx context.Context, id string) error
}

type service struct {
	logger              logger.Logger
	inventoryRepository inventoryRepository.Repository
	usageUnitRepository usageUnitRepository.Repository
}

func New(
	logger logger.Logger,
	inventoryRepository inventoryRepository.Repository,
	usageUnitRepository usageUnitRepository.Repository,
) Service {
	return &service{
		logger:              logger,
		inventoryRepository: inventoryRepository,
		usageUnitRepository: usageUnitRepository,
	}
}

func (s *service) Create(ctx context.Context, req *inventoryModule.Request) (*handling.ResponseItem[*inventoryModule.Request], error) {
	usageUnit, err := s.usageUnitRepository.FindByCode(ctx, req.PurchaseUnit.Code)
	if err != nil {
		if !errors.Is(err, define.ErrRecordNotFound) {
			s.logger.Error(err)
		}

		return nil, handling.ThrowErr(err)
	}

	req.PurchaseUnit.Name = usageUnit.Name

	err = s.inventoryRepository.Create(ctx, req)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*inventoryModule.Request]{
		Item: req,
	}, nil
}

func (s *service) Find(ctx context.Context) (*handling.ResponseItems[inventoryModule.Response], error) {
	total, err := s.inventoryRepository.Count(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	inventories, err := s.inventoryRepository.Find(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItems[inventoryModule.Response]{
		Items: inventories,
		Meta: handling.MetaResponse{
			Total: total,
		},
	}, nil
}

func (s *service) FindByID(ctx context.Context, id string) (*handling.ResponseItem[*inventoryModule.Response], error) {
	inventory, err := s.inventoryRepository.FindByID(ctx, id)
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
	_, err := s.inventoryRepository.FindByID(ctx, id)
	if err != nil {
		if !errors.Is(err, define.ErrRecordNotFound) {
			s.logger.Error(err)
		}

		return nil, handling.ThrowErr(err)
	}

	usageUnit, err := s.usageUnitRepository.FindByCode(ctx, req.PurchaseUnit.Code)
	if err != nil {
		if !errors.Is(err, define.ErrRecordNotFound) {
			s.logger.Error(err)
		}

		return nil, handling.ThrowErr(err)
	}

	req.PurchaseUnit.Name = usageUnit.Name

	err = s.inventoryRepository.UpdateByID(ctx, id, req)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*inventoryModule.Request]{
		Item: req,
	}, nil
}

func (s *service) DeleteByID(ctx context.Context, id string) error {
	_, err := s.inventoryRepository.FindByID(ctx, id)
	if err != nil {
		if !errors.Is(err, define.ErrRecordNotFound) {
			s.logger.Error(err)
		}

		return handling.ThrowErr(err)
	}

	err = s.inventoryRepository.DeleteByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return handling.ThrowErr(err)
	}

	return nil
}
