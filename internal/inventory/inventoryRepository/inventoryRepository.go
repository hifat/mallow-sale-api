package inventoryRepository

import (
	"context"

	"github.com/hifat/cost-calculator-api/internal/inventory"
)

type IInventoryRepository interface {
	Create(ctx context.Context, req inventory.Inventory) (string, error)
	Find(ctx context.Context) ([]inventory.Inventory, error)
}
