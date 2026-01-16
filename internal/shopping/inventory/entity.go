package shoppingInventoryModule

import (
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
)

type Entity struct {
	utilsModule.Base `bson:"inline"`

	InventoryID   string `bson:"inventory_id" json:"inventoryID"`
	InventoryName string `bson:"inventory_name" json:"inventoryName"`
	SupplierID    string `bson:"supplier_id" json:"supplierID"`
	SupplierName  string `bson:"supplier_name" json:"supplierName"`
}
