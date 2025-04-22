package inventoryRepository

import (
	"context"

	"github.com/hifat/cost-calculator-api/internal/inventory"
)

type IInventoryRepository interface {
	Create(ctx context.Context, req inventory.InventoryReq) (string, error)
	Find(ctx context.Context) ([]inventory.Inventory, error)
	FindByID(ctx context.Context, id string) (*inventory.Inventory, error)
	FindInID(ctx context.Context, ids []string) ([]inventory.Inventory, error)
	Update(ctx context.Context, id string, req inventory.InventoryReq) error
	Delete(ctx context.Context, id string) error
}
