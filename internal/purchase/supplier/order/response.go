package purchaseSupplierOrderModule

import (
	"time"

	purchaseStatusModule "github.com/hifat/mallow-sale-api/internal/purchase/status"
)

type Response struct {
	ID                 string                                      `json:"id"`
	PurchaseSupplierID string                                      `json:"purchaseSupplierID"`
	InventoryID        string                                      `json:"inventoryID"`
	InventoryName      string                                      `json:"inventoryName"`
	Quantity           float64                                     `json:"quantity"`
	UsageUnitCode      string                                      `json:"usageUnitCode"`
	StatusCode         purchaseStatusModule.EnumPurchaseStatusCode `json:"statusCode"`
	CreatedAt          time.Time                                   `json:"createdAt"`
	UpdatedAt          time.Time                                   `json:"updatedAt"`
}
