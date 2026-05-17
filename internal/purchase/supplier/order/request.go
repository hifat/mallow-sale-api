package purchaseSupplierOrderModule

import (
	purchaseStatusModule "github.com/hifat/mallow-sale-api/internal/purchase/status"
)

type CreateOrderRequest struct {
	InventoryID   string                                      `json:"inventory_id" binding:"required"`
	InventoryName string                                      `json:"inventory_name" binding:"required"`
	Quantity      float64                                     `json:"quantity" binding:"required"`
	UsageUniCode  string                                      `json:"usage_unit" binding:"required"`
	UnitPrice     float64                                     `json:"unit_price" binding:"required"`
	TotalPrice    float64                                     `json:"total_price" binding:"required"`
	Status        purchaseStatusModule.EnumPurchaseStatusCode `json:"status" binding:"required"`

	PurchaseSupplierID string `json:"-"`
}
