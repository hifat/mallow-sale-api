package usageUnitService

import (
	"context"

	core "github.com/hifat/goroger-core"
	"github.com/hifat/mallow-sale-api/internal/usageUnit"
	"github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitRepository"
	"github.com/hifat/mallow-sale-api/pkg/throw"
)

type IUsageUnitService interface {
	FindIn(ctx context.Context, filter usageUnit.FilterReq) ([]usageUnit.UsageUnit, error)
}

type usageUnitService struct {
	logger        core.Logger
	usageUnitRepo usageUnitRepository.IUsageUnitRepository
}

func New(logger core.Logger, usageUnitRepo usageUnitRepository.IUsageUnitRepository) IUsageUnitService {
	return &usageUnitService{
		logger,
		usageUnitRepo,
	}
}

func (s *usageUnitService) FindIn(ctx context.Context, filter usageUnit.FilterReq) ([]usageUnit.UsageUnit, error) {
	usageUnits, err := s.usageUnitRepo.FindInCodes(ctx, filter.Codes)
	if err != nil {
		s.logger.Error(err)
		return nil, throw.InternalServerErr(err)
	}

	return usageUnits, nil
}
