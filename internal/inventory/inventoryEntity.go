package inventory

import (
	"github.com/hifat/cost-calculator-api/internal/entity"
	"github.com/hifat/cost-calculator-api/internal/usageUnit"
)

type (
	Inventory struct {
		entity.Base `bson:"inline"`

		Name             string               `bson:"name"`
		PurchasePrice    float64              `bson:"purchase_price"`
		PurchaseQuantity float64              `bson:"purchase_quantity"`
		PurchaseUnit     *usageUnit.UsageUnit `bson:"purchase_unit"`
		YieldPercentage  float64              `bson:"yield_percentage"`
	}
)

func (m *Inventory) DocName() string {
	return "inventories"
}
