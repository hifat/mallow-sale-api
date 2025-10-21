package shoppingModule

import usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"

type Request struct {
	InventoryID      string                 `fake:"{inventoryID}" json:"inventoryID"`
	PurchaseQuantity float64                `fake:"{float64}" json:"purchaseQuantity"`
	PurchaseUnit     usageUnitModule.Entity `json:"purchaseUnit"`
}

type UpdateIsComplete struct {
	IsComplete bool `json:"is_complete"`
}
