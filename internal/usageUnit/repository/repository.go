package usageUnitRepository

import (
	"context"

	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
)

//go:generate mockgen -source=./repository.go -destination=./mock/repository.go -package=mockUsageUnitRepository
type IRepository interface {
	FindByCode(ctx context.Context, code string) (*usageUnitModule.Prototype, error)
	FindInCodes(ctx context.Context, codes []string) ([]usageUnitModule.Prototype, error)
}
