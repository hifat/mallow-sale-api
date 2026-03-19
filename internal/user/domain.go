package userModule

import (
	"context"
)

//go:generate mockgen -source=./domain.go -destination=./repository/mock/repository.go -package=mockUsernameRepository
type IRepository interface {
	FindByUsername(ctx context.Context, username string) (*Response, error)
	FindByEmail(ctx context.Context, email string) (*Response, error)
}
