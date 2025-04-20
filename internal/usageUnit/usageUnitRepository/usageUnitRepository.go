package usageUnitRepository

import (
	"context"

	"github.com/hifat/cost-calculator-api/internal/usageUnit"
)

type IUsageUnitRepository interface {
	FindInCodes(ctx context.Context, codes []string) ([]usageUnit.UsageUnit, error)
}
