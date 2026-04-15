package pricePresetModule

import (
	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"time"
)

type Price struct {
	ID        string    `bson:"id" json:"id"`
	StockID   string    `bson:"stock_id" json:"stockID"`
	Price     float64   `bson:"price" json:"price"`
	CreatedAt time.Time `bson:"created_at" json:"createdAt"`
}

type Entity struct {
	utilsModule.Base `bson:"inline"`

	InventoryID string                  `bson:"inventory_id" json:"inventoryID"`
	Inventory   *inventoryModule.Entity `bson:"inventory" json:"inventory"`

	Prices []Price `bson:"prices" json:"prices"`
}
