package settingService

import (
	"context"

	settingModule "github.com/hifat/mallow-sale-api/internal/settings"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type service struct {
	settingRepo settingModule.IRepository
	logger      logger.ILogger
}

func New(settingRepo settingModule.IRepository, logger logger.ILogger) settingModule.IService {
	return &service{settingRepo: settingRepo, logger: logger}
}

func (s *service) Update(ctx context.Context, req *settingModule.Request) error {
	err := s.settingRepo.Update(ctx, req)
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
