package purchasePresetModule

type InventoryRequest struct {
	No               uint   `bson:"no" json:"no"`
	Name             string `bson:"name" json:"name"`
	PurchaseUnitCode string `bson:"purchase_unit_code" json:"purchaseUnitCode"`
}

type Request struct {
	SupplierID   string             `json:"supplierID"`
	SupplierName string             `json:"supplierName"`
	Inventories  []InventoryRequest `json:"inventories"`
}
