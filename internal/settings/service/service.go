package settingService

import (
	settingModule "github.com/hifat/mallow-sale-api/internal/settings"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type Service interface {
	Update(costPercentage float32) error
	Get() (*settingModule.Entity, error)
}

type service struct {
	repo   settingModule.Repository
	logger logger.Logger
}

func New(repo settingModule.Repository, logger logger.Logger) Service {
	return &service{repo: repo, logger: logger}
}

func (s *service) Update(costPercentage float32) error {
	err := s.repo.Update(costPercentage)
	if err != nil {
		s.logger.Error(err)
	}

	return err
}

func (s *service) Get() (*settingModule.Entity, error) {
	settings, err := s.repo.Get()
	if err != nil {
		s.logger.Error(err)
	}

	return settings, err
}
