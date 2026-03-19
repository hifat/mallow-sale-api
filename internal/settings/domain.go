package settingModule

import (
	"context"

	"github.com/hifat/mallow-sale-api/pkg/handling"
)

//go:generate mockgen -source=./domain.go -destination=./repository/mock/repository.go -package=mockSettingRepository
type IRepository interface {
	Find(ctx context.Context) (*Response, error)
	Update(ctx context.Context, req *Request) error
}

type IService interface {
	Update(ctx context.Context, req *Request) error
	Find(ctx context.Context) (*handling.ResponseItem[*Response], error)
}
