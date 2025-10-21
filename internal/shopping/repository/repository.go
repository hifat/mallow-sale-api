package shoppingRepository

import shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"

//go:generate mockgen -source=./repository.go -destination=./mock/repository.go -package=mockShoppingRepository
type IRepository interface {
	Create(req *shoppingModule.Request) error
	UpdateIsComplete(req *shoppingModule.UpdateIsComplete) error
	Delete(id string) error
}
