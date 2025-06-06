package inventory

import (
	"github.com/hifat/mallow-sale-api/internal/entity"
	"github.com/hifat/mallow-sale-api/internal/usageUnit"
)

type (
	Inventory struct {
		entity.Base `bson:"inline"`

		Name             string                   `fake:"{name}" bson:"name"`
		PurchasePrice    float32                  `fake:"{float32}" bson:"purchase_price"`
		PurchaseQuantity float32                  `fake:"{float32}" bson:"purchase_quantity"`
		PurchaseUnit     usageUnit.UsageUnitEmbed `bson:"purchase_unit"`
		YieldPercentage  float32                  `fake:"{float32}" bson:"yield_percentage"`
		Remark           string                   `fake:"{sentence}" bson:"remark"`
	}
)

func (m *Inventory) Doc() string {
	return "inventories"
}
