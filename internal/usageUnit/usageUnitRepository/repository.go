package usageUnitRepository

import (
	"context"

	"github.com/hifat/mallow-sale-api/internal/usageUnit"
)

type IUsageUnitRepository interface {
	FindInCodes(ctx context.Context, codes []string) ([]usageUnit.UsageUnit, error)
}
