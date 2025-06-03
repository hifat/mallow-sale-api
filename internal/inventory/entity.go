package inventory

import (
	"github.com/hifat/mallow-sale-api/internal/entity"
	"github.com/hifat/mallow-sale-api/internal/usageUnit"
)

type (
	Inventory struct {
		entity.Base `bson:"inline"`

		Name             string                   `bson:"name"`
		PurchasePrice    float32                  `bson:"purchase_price"`
		PurchaseQuantity float32                  `bson:"purchase_quantity"`
		PurchaseUnit     usageUnit.UsageUnitEmbed `bson:"purchase_unit"`
		YieldPercentage  float32                  `bson:"yield_percentage"`
		Remark           string                   `bson:"remark"`
	}
)

func (m *Inventory) Doc() string {
	return "inventories"
}
