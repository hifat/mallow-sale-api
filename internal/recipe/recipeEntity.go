package recipe

import (
	"github.com/hifat/cost-calculator-api/internal/entity"
	"github.com/hifat/cost-calculator-api/internal/inventory"
	"github.com/hifat/cost-calculator-api/internal/usageUnit"
)

type (
	RecipeInventory struct {
		entity.Base   `bson:"inline"`
		UsageQuantity float64 `bson:"usage_quantity"`
		Remark        string  `bson:"remark"`

		InventoryID string               `bson:"inventory_id"`
		Inventory   *inventory.Inventory `bson:"inventory"`

		UsageUnit *usageUnit.UsageUnitEmbed `bson:"usage_unit"`
	}

	Recipe struct {
		entity.Base `bson:"inline"`
		Name        string            `bson:"name"`
		Inventories []RecipeInventory `bson:"inventories"`
	}
)

func (m *Recipe) DocName() string {
	return "inventories"
}
