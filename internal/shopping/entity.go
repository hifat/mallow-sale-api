package shoppingModule

import (
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
)

type Entity struct {
	utilsModule.Base

	InventoryName    string                 `fake:"{name}" bson:"name"`
	PurchaseQuantity float64                `fake:"{float64}" bson:"purchase_quantity"`
	PurchaseUnit     usageUnitModule.Entity `bson:"purchase_unit"`
	IsComplete       bool                   `bson:"is_complete"`
}
