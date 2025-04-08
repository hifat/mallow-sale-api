package inventoryRepository

import (
	"context"

	"github.com/hifat/cost-calculator-api/internal/inventory"
)

type IInventoryRepository interface {
	Create(ctx context.Context, req inventory.Inventory) (string, error)
	Find(ctx context.Context) ([]inventory.Inventory, error)
	FindByID(ctx context.Context, id string) (*inventory.Inventory, error)
	Update(ctx context.Context, id string, req inventory.Inventory) error
	Delete(ctx context.Context, id string) error
}
