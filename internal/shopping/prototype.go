package shoppingModule

import usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"

type PrototypeInventory struct {
	OrderNo          uint                   `json:"orderNo"`
	InventoryID      string                 `json:"inventoryID"`
	InventoryName    string                 `json:"inventoryName"`
	PurchaseQuantity float64                `json:"purchaseQuantity"`
	PurchaseUnit     usageUnitModule.Entity `json:"purchaseUnit"`
	Status           InventoryStatus        `json:"status"`
}
