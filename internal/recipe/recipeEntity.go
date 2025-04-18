package recipe

import (
	"github.com/hifat/cost-calculator-api/internal/entity"
	"github.com/hifat/cost-calculator-api/internal/inventory"
	"github.com/hifat/cost-calculator-api/internal/usageUnit"
)

type RecipeInventory struct {
	entity.Base   `bson:"inline"`
	UsageQuantity float64 `json:"usageQuantity" bson:"usage_quantity"`
	Remark        string  `json:"remark" bson:"remark"`

	UsageUnit *usageUnit.UsageUnitEmbed `json:"usageUnit" bson:"usage_unit"`

	InventoryID string               `json:"inventoryID" bson:"inventory_id"`
	Inventory   *inventory.Inventory `json:"inventory" bson:"inventory,omitempty"`
}

type Recipe struct {
	entity.Base `bson:"inline"`
	Name        string            `json:"name" bson:"name"`
	Inventories []RecipeInventory `json:"inventories" bson:"inventories"`
}

func (m *Recipe) DocName() string {
	return "recipes"
}
