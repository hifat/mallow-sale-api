package purchaseSupplierOrderModule

import (
	purchaseStatusModule "github.com/hifat/mallow-sale-api/internal/purchase/status"
)

type CreateOrderRequest struct {
	InventoryID   string                                      `json:"inventoryID" binding:"required"`
	InventoryName string                                      `json:"inventoryName" binding:"required"`
	Quantity      float64                                     `json:"quantity" binding:"required"`
	UsageUnitCode string                                      `json:"usageUnitCode" binding:"required"`
	UnitPrice     float64                                     `json:"unitPrice" binding:"required"`
	TotalPrice    float64                                     `json:"totalPrice" binding:"required"`
	StatusCode    purchaseStatusModule.EnumPurchaseStatusCode `json:"statusCode" binding:"required"`

	PurchaseSupplierID string `json:"-"`
}
