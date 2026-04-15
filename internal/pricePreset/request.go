package pricePresetModule

type Request struct {
	InventoryID string  `validate:"required" json:"inventoryID"`
	StockID     string  `validate:"required" json:"stockID"`
	Price       float64 `validate:"required" json:"price"`
}
