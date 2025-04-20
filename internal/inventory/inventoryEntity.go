package inventory

import (
	"github.com/hifat/cost-calculator-api/internal/entity"
	"github.com/hifat/cost-calculator-api/internal/usageUnit"
)

type (
	Inventory struct {
		entity.Base `bson:"inline"`

		Name             string                    `bson:"name"`
		PurchasePrice    float64                   `bson:"purchase_price"`
		PurchaseQuantity float64                   `bson:"purchase_quantity"`
		PurchaseUnit     *usageUnit.UsageUnitEmbed `bson:"purchase_unit"`
		YieldPercentage  float64                   `bson:"yield_percentage"`
		Remark           string                    `bson:"remark"`
	}
)

func (m *Inventory) DocName() string {
	return "inventories"
}
