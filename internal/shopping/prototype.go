package shoppingModule

import usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"

type PrototypeStatus struct {
	Code EnumCodeShoppingStatusType `bson:"code" json:"code"`
	Name string                     `bson:"name" json:"name"`
}

type PrototypeInventoryStatus struct {
	Code EnumCodeInventoryStatusType `bson:"code" json:"code"`
	Name string                      `bson:"name" json:"name"`
}

type PrototypeInventory struct {
	OrderNo          uint                     `json:"orderNo"`
	InventoryID      string                   `json:"inventoryID"`
	InventoryName    string                   `json:"inventoryName"`
	PurchaseQuantity float64                  `json:"purchaseQuantity"`
	PurchaseUnit     usageUnitModule.Entity   `json:"purchaseUnit"`
	Status           PrototypeInventoryStatus `json:"status"`
}
