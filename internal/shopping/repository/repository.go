package shoppingRepository

import (
	"context"

	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
)

//go:generate mockgen -source=./repository.go -destination=./mock/repository.go -package=mockShoppingRepository
type IRepository interface {
	Create(ctx context.Context, req *shoppingModule.Request) error
	UpdateIsComplete(ctx context.Context, req *shoppingModule.UpdateIsComplete) error
	Delete(ctx context.Context, id string) error
}
