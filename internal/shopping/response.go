package shoppingModule

import "time"

type Response struct {
	ID           string               `json:"id"`
	SupplierID   string               `json:"supplierID"`
	SupplierName string               `json:"supplierName"`
	Status       PrototypeStatus      `json:"status"`
	CreatedAt    time.Time            `json:"createdAt"`
	UpdatedAt    time.Time            `json:"updatedAt"`
	Inventories  []PrototypeInventory `json:"inventories"`
}

type ResReceiptReader struct {
	InventoryID      string  `json:"inventoryID"` // Make vector db solution for search
	Name             string  `json:"name"`
	NameEdited       string  `json:"nameEdited"`
	PurchasePrice    float64 `json:"purchasePrice"`
	PurchaseQuantity float64 `json:"purchaseQuantity"`
	Remark           string  `json:"remark"`
}

/* --------------------------- Shopping Inventory --------------------------- */

type InventoryResponse struct {
	ID            string `bson:"_id" json:"id"`
	InventoryID   string `bson:"inventory_id" json:"inventoryID"`
	InventoryName string `bson:"inventory_name" json:"inventoryName"`
	SupplierID    string `bson:"supplier_id" json:"supplierID"`
	SupplierName  string `bson:"supplier_name" json:"supplierName"`
}
