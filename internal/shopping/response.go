package shoppingModule

import (
	"time"
)

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

type ResInventory struct {
	ID            string `bson:"id" json:"id"`
	InventoryID   string `bson:"inventoryID" json:"inventoryID"`
	InventoryName string `bson:"inventoryName" json:"inventoryName"`
	UsageUnitCode string `bson:"usageUnitCode" json:"usageUnitCode"`
}

type ResShoppingInventory struct {
	ID           string         `bson:"_id" json:"id"`
	SupplierID   string         `bson:"supplierID" json:"supplierID"`
	SupplierName string         `bson:"supplierName" json:"supplierName"`
	Inventories  []ResInventory `bson:"inventories" json:"inventories"`
}
