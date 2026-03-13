package usageUnitModule

import (
	"context"
)

//go:generate mockgen -source=./repository.go -destination=./mock/repository/repository.go -package=mockUsageUnitRepository
type IRepository interface {
	FindByCode(ctx context.Context, code string) (*Prototype, error)
	FindInCodes(ctx context.Context, codes []string) ([]Prototype, error)
}
