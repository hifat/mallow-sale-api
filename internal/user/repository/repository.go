package userRepository

import (
	"context"

	userModule "github.com/hifat/mallow-sale-api/internal/user"
)

//go:generate mockgen -source=./repository.go -destination=./mock/repository.go -package=mockUsernameRepository
type IRepository interface {
	FindByUsername(ctx context.Context, username string) (*userModule.Response, error)
}
