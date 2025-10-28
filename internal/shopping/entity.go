package shoppingModule

import (
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
)

type Entity struct {
	utilsModule.Base `bson:"inline"`

	OrderNo          uint                   `fake:"{uintrange:0,100}" bson:"order_no" json:"orderNo"`
	Name             string                 `fake:"{name}" bson:"name" json:"name"`
	PurchaseQuantity float64                `fake:"{float64}" bson:"purchase_quantity" json:"purchaseQuantity"`
	PurchaseUnit     usageUnitModule.Entity `bson:"purchase_unit" json:"purchaseUnit"`
	IsComplete       bool                   `bson:"is_complete" json:"isComplete"`
}
