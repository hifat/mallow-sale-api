package stockModule

import (
	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	supplierModule "github.com/hifat/mallow-sale-api/internal/supplier"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
)

type Entity struct {
	utilsModule.Base `bson:"inline"`

	InventoryID string                  `bson:"inventory_id" json:"inventoryID"`
	Inventory   *inventoryModule.Entity `bson:"inventory" json:"inventory"`

	SupplierID string                 `bson:"supplier_id" json:"supplierID"`
	Supplier   *supplierModule.Entity `bson:"supplier" json:"supplier"`

	PurchasePrice    float64                `fake:"{float64}" bson:"purchase_price" json:"purchasePrice"`
	PurchaseQuantity float64                `fake:"{float64}" bson:"purchase_quantity" json:"purchaseQuantity"`
	PurchaseUnit     usageUnitModule.Entity `bson:"purchase_unit" json:"purchaseUnit"`
	Remark           string                 `fake:"{sentence}" bson:"remark" json:"remark"`
}
