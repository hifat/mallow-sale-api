package shoppingInventoryModule

type Request struct {
	InventoryID   string `validate:"required" json:"inventoryID"`
	InventoryName string `validate:"required" json:"inventoryName"`
	SupplierID    string `validate:"required" json:"supplierID"`
	SupplierName  string `validate:"required" json:"supplierName"`
}
