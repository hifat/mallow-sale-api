package purchasePresetModule

type InventoryRequest struct {
	ID               string `json:"id"`
	No               uint   `json:"no"`
	Name             string `json:"name"`
	PurchaseUnitCode string `json:"purchaseUnitCode"`
}

type Request struct {
	SupplierID   string             `json:"supplierID"`
	SupplierName string             `json:"supplierName"`
	Inventories  []InventoryRequest `json:"inventories"`
}
