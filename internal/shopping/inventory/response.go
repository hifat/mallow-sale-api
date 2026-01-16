package shoppingInventoryModule

type Response struct {
	ID            string `bson:"_id" json:"id"`
	InventoryID   string `bson:"inventory_id" json:"inventoryID"`
	InventoryName string `bson:"inventory_name" json:"inventoryName"`
	SupplierID    string `bson:"supplier_id" json:"supplierID"`
	SupplierName  string `bson:"supplier_name" json:"supplierName"`
}
