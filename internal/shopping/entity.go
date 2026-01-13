package shoppingModule

import (
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
)

type InventoryStatus struct {
	Code EnumCodeInventoryStatusType `bson:"code" json:"code"`
	Name string                      `bson:"name" json:"name"`
}

type Inventory struct {
	OrderNo          uint                   `fake:"{uintrange:0,100}" bson:"order_no" json:"orderNo"`
	InventoryID      string                 `bson:"inventory_id" json:"inventoryID"`
	InventoryName    string                 `bson:"inventory_name" json:"inventoryName"`
	PurchaseQuantity float64                `fake:"{float64}" bson:"purchase_quantity" json:"purchaseQuantity"`
	PurchaseUnit     usageUnitModule.Entity `bson:"purchase_unit" json:"purchaseUnit"`
	Status           InventoryStatus        `bson:"status" json:"status"`
}

type Entity struct {
	utilsModule.Base `bson:"inline"`

	SupplierID   string      `bson:"supplier_id" json:"supplierID"`
	SupplierName string      `bson:"supplier_name" json:"supplierName"`
	Inventories  []Inventory `bson:"inventories" json:"inventories"`
}
