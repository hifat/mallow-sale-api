package purchasePresetModule

import "time"

type InventoryResponse struct {
	ID               string `json:"id"`
	No               uint   `json:"no"`
	Name             string `json:"name"`
	PurchaseUnitCode string `json:"purchaseUnitCode"`
}

type Response struct {
	ID           string              `json:"id"`
	SupplierID   string              `json:"supplierID"`
	SupplierName string              `json:"supplierName"`
	Inventories  []InventoryResponse `json:"inventories"`
	CreatedAt    *time.Time          `json:"createdAt"`
	UpdatedAt    *time.Time          `json:"updatedAt"`
}
