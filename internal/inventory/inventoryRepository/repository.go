package inventoryRepository

import (
	"context"

	"github.com/hifat/mallow-sale-api/internal/inventory"
)

//go:generate mockgen -source=./repository.go -destination=./mock/repository.go -package=mockInventoryRepository
type IInventoryRepository interface {
	Create(ctx context.Context, req inventory.InventoryReq) (string, error)
	Find(ctx context.Context) ([]inventory.Inventory, error)
	FindByID(ctx context.Context, id string) (*inventory.Inventory, error)
	FindIn(ctx context.Context, filter inventory.FilterReq) ([]inventory.Inventory, error)
	Update(ctx context.Context, id string, req inventory.InventoryReq) error
	Delete(ctx context.Context, id string) error
}
