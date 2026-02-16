package settingModule

import (
	"context"

	"github.com/hifat/mallow-sale-api/pkg/handling"
)

//go:generate mockgen -source=./repository.go -destination=./mock/repository/repository.go -package=mockSettingRepository
type IRepository interface {
	Find(ctx context.Context) (*Response, error)
	Update(ctx context.Context, costPercentage float32) error
}

type IService interface {
	Update(ctx context.Context, costPercentage float32) error
	Find(ctx context.Context) (*handling.ResponseItem[*Response], error)
}
