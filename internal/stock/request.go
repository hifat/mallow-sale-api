package stockModule

import usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"

type Request struct {
	InventoryID      string                       `validate:"required" json:"inventoryID"`
	SupplierID       string                       `validate:"required" json:"supplierID"`
	PurchasePrice    float64                      `validate:"required" json:"purchasePrice"`
	PurchaseQuantity float64                      `validate:"required" json:"purchaseQuantity"`
	PurchaseUnit     usageUnitModule.UsageUnitReq `validate:"required" json:"purchaseUnit"`
	Remark           string                       `json:"remark"`
}
