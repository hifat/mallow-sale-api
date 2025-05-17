package usageUnitServiceUtils

import (
	"context"

	"github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitRepository"
)

type IUsageUnitServiceUtils interface {
	MapUsageUnitName(ctx context.Context, codes []string) (map[string]string, error)
}

type usageUnitServiceUtils struct {
	usageUnitRepo usageUnitRepository.IUsageUnitRepository
}

func New(usageUnitRepo usageUnitRepository.IUsageUnitRepository) IUsageUnitServiceUtils {
	return &usageUnitServiceUtils{
		usageUnitRepo,
	}
}

func (s *usageUnitServiceUtils) MapUsageUnitName(ctx context.Context, codes []string) (map[string]string, error) {
	_usageUnits, err := s.usageUnitRepo.FindInCodes(ctx, codes)
	if err != nil {
		return nil, err
	}

	unitCodeMap := make(map[string]string)
	for _, usageUnit := range _usageUnits {
		unitCodeMap[usageUnit.Code] = usageUnit.Name
	}

	return unitCodeMap, nil
}
