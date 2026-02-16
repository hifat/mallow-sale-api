package userModule

import (
	"context"
)

//go:generate mockgen -source=./repository.go -destination=./mock/repository/repository.go -package=mockUsernameRepository
type IRepository interface {
	FindByUsername(ctx context.Context, username string) (*Response, error)
}
