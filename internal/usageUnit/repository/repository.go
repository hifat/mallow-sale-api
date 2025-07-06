package usageUnitRepository

import (
	"context"

	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
)

type Repository interface {
	FindByCode(ctx context.Context, code string) (*usageUnitModule.Prototype, error)
	FindInCodes(ctx context.Context, codes []string) ([]usageUnitModule.Prototype, error)
}
