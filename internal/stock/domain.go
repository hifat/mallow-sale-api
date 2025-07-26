package stockModule

import (
	"time"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	supplierModule "github.com/hifat/mallow-sale-api/internal/supplier"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
)

type Request struct {
	InventoryID      string                       `validate:"required" json:"inventoryID"`
	SupplierID       string                       `validate:"required" json:"supplierID"`
	PurchasePrice    float32                      `validate:"required" json:"purchasePrice"`
	PurchaseQuantity float32                      `validate:"required" json:"purchaseQuantity"`
	PurchaseUnit     usageUnitModule.UsageUnitReq `validate:"required" json:"purchaseUnit"`
	Remark           string                       `json:"remark"`
}

type Prototype struct {
	ID               string                     `json:"id"`
	InventoryID      string                     `json:"inventoryID"`
	Inventory        *inventoryModule.Prototype `json:"inventory"`
	SupplierID       string                     `json:"supplierID"`
	Supplier         *supplierModule.Prototype  `json:"supplier"`
	PurchasePrice    float32                    `json:"purchasePrice"`
	PurchaseQuantity float32                    `json:"purchaseQuantity"`
	PurchaseUnit     usageUnitModule.Prototype  `json:"purchaseUnit"`
	Remark           string                     `json:"remark"`
	CreatedAt        *time.Time                 `json:"createdAt"`
	UpdatedAt        *time.Time                 `json:"updatedAt"`
}

type Response struct {
	Prototype
}
