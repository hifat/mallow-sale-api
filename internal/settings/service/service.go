package settingService

import (
	"context"

	settingModule "github.com/hifat/mallow-sale-api/internal/settings"
	settingRepository "github.com/hifat/mallow-sale-api/internal/settings/repository"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type IService interface {
	Update(ctx context.Context, costPercentage float32) error
	Find(ctx context.Context) (*handling.ResponseItem[*settingModule.Response], error)
}

type service struct {
	settingRepo settingRepository.IRepository
	logger      logger.ILogger
}

func New(settingRepo settingRepository.IRepository, logger logger.ILogger) IService {
	return &service{settingRepo: settingRepo, logger: logger}
}

func (s *service) Update(ctx context.Context, costPercentage float32) error {
	err := s.settingRepo.Update(ctx, costPercentage)
	if err != nil {
		s.logger.Error(err)
		return handling.ThrowErr(err)
	}

	return nil
}

func (s *service) Find(ctx context.Context) (*handling.ResponseItem[*settingModule.Response], error) {
	setting, err := s.settingRepo.Find(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*settingModule.Response]{
		Item: setting,
	}, nil
}
