package usageUnitHelper

import (
	"context"

	usageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type IHelper interface {
	GetNameByCode(ctx context.Context, findInCodes []string) (func(code string) (name string), error)
}

type helper struct {
	logger        logger.ILogger
	usageUnitRepo usageUnitRepository.IRepository
}

func New(
	logger logger.ILogger,
	usageUnitRepo usageUnitRepository.IRepository,
) IHelper {
	return &helper{
		logger:        logger,
		usageUnitRepo: usageUnitRepo,
	}
}

func (h *helper) GetNameByCode(ctx context.Context, findInCodes []string) (func(code string) (name string), error) {
	usageUnits, err := h.usageUnitRepo.FindInCodes(ctx, findInCodes)
	if err != nil {
		return nil, err
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
