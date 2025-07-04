package inventoryModule

import (
	entityModule "github.com/hifat/mallow-sale-api/internal/entity"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
)

type Entity struct {
	entityModule.Base `bson:"inline"`

	Name             string                 `fake:"{name}" bson:"name"`
	PurchasePrice    float32                `fake:"{float32}" bson:"purchase_price"`
	PurchaseQuantity float32                `fake:"{float32}" bson:"purchase_quantity"`
	PurchaseUnit     usageUnitModule.Entity `bson:"purchase_unit"`
	YieldPercentage  float32                `fake:"{float32}" bson:"yield_percentage"`
	Remark           string                 `fake:"{sentence}" bson:"remark"`
}
