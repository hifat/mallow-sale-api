package stockModule

import (
	"time"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	supplierModule "github.com/hifat/mallow-sale-api/internal/supplier"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
)

type Prototype struct {
	ID               string                     `json:"id"`
	InventoryID      string                     `json:"inventoryID"`
	Inventory        *inventoryModule.Prototype `json:"inventory"`
	SupplierID       string                     `json:"supplierID"`
	Supplier         *supplierModule.Prototype  `json:"supplier"`
	PurchasePrice    float64                    `json:"purchasePrice"`
	PurchaseQuantity float64                    `json:"purchaseQuantity"`
	PurchaseUnit     usageUnitModule.Prototype  `json:"purchaseUnit"`
	Remark           string                     `json:"remark"`
	CreatedAt        *time.Time                 `json:"createdAt"`
	UpdatedAt        *time.Time                 `json:"updatedAt"`
}
