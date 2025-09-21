package supplierService

import (
	"context"

	supplierModule "github.com/hifat/mallow-sale-api/internal/supplier"
	supplierRepository "github.com/hifat/mallow-sale-api/internal/supplier/repository"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type IService interface {
	Create(ctx context.Context, req *supplierModule.Request) (*handling.ResponseItem[*supplierModule.Request], error)
	Find(ctx context.Context, query *utilsModule.QueryReq) (*handling.ResponseItems[supplierModule.Response], error)
	FindByID(ctx context.Context, id string) (*handling.ResponseItem[*supplierModule.Response], error)
	UpdateByID(ctx context.Context, id string, req *supplierModule.Request) (*handling.ResponseItem[*supplierModule.Request], error)
	DeleteByID(ctx context.Context, id string) error
}

type service struct {
	supplierRepository supplierRepository.IRepository
	logger             logger.Logger
}

func New(
	supplierRepository supplierRepository.IRepository,
	logger logger.Logger,
) IService {
	return &service{
		supplierRepository: supplierRepository,
		logger:             logger,
	}
}

func (s *service) Create(ctx context.Context, req *supplierModule.Request) (*handling.ResponseItem[*supplierModule.Request], error) {
	err := s.supplierRepository.Create(ctx, req)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*supplierModule.Request]{Item: req}, nil
}

func (s *service) Find(ctx context.Context, query *utilsModule.QueryReq) (*handling.ResponseItems[supplierModule.Response], error) {
	count, err := s.supplierRepository.Count(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	suppliers, err := s.supplierRepository.Find(ctx, query)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItems[supplierModule.Response]{
		Items: suppliers,
		Meta:  handling.MetaResponse{Total: count},
	}, nil
}

func (s *service) FindByID(ctx context.Context, id string) (*handling.ResponseItem[*supplierModule.Response], error) {
	supplier, err := s.supplierRepository.FindByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*supplierModule.Response]{Item: supplier}, nil
}

func (s *service) UpdateByID(ctx context.Context, id string, req *supplierModule.Request) (*handling.ResponseItem[*supplierModule.Request], error) {
	err := s.supplierRepository.UpdateByID(ctx, id, req)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*supplierModule.Request]{Item: req}, nil
}

func (s *service) DeleteByID(ctx context.Context, id string) error {
	err := s.supplierRepository.DeleteByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return handling.ThrowErr(err)
	}

	return nil
}
