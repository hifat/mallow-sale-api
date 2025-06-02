package usageUnitRepository

import (
	"context"

	"github.com/hifat/mallow-sale-api/internal/usageUnit"
)

//go:generate mockgen -source=./repository.go -destination=./mock/repository.go -package=mockUsageUnitRepository
type IUsageUnitRepository interface {
	FindInCodes(ctx context.Context, codes []string) ([]usageUnit.UsageUnit, error)
}
