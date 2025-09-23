package settingRepository

import (
	"context"

	settingModule "github.com/hifat/mallow-sale-api/internal/settings"
)

//go:generate mockgen -source=./repository.go -destination=./mock/repository.go -package=mockSettingRepository
type IRepository interface {
	Find(ctx context.Context) (*settingModule.Response, error)
	Update(ctx context.Context, costPercentage float32) error
}
