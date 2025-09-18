package usageUnitHelper

import (
	"context"

	usageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type Helper interface {
	// Returns a function that maps usage unit codes to their names
	GetNameByCode(ctx context.Context, findInCodes []string) (func(code string) (name string), error)
}

type helper struct {
	logger              logger.Logger
	usageUnitRepository usageUnitRepository.IRepository
}

func New(
	logger logger.Logger,
	usageUnitRepository usageUnitRepository.IRepository,
) Helper {
	return &helper{
		logger:              logger,
		usageUnitRepository: usageUnitRepository,
	}
}

func (h *helper) GetNameByCode(ctx context.Context, findInCodes []string) (func(code string) (name string), error) {
	usageUnits, err := h.usageUnitRepository.FindInCodes(ctx, findInCodes)
	if err != nil {
		return nil, handling.ThrowErr(err)
	}

	mapUsageUnit := make(map[string]string)
	for _, usageUnit := range usageUnits {
		mapUsageUnit[usageUnit.Code] = usageUnit.Name
	}

	return func(code string) (name string) {
		name, ok := mapUsageUnit[code]
		if !ok {
			return ""
		}

		return name
	}, nil
}
