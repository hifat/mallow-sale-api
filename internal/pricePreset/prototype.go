package pricePresetModule

import (
	"time"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
)

type PricePrototype struct {
	ID        string     `json:"id"`
	StockID   string     `json:"stockID"`
	Price     float64    `json:"price"`
	CreatedAt time.Time  `json:"createdAt"`
}

type Prototype struct {
	ID          string                     `json:"id"`
	InventoryID string                     `json:"inventoryID"`
	Inventory   *inventoryModule.Prototype `json:"inventory"`
	Prices      []PricePrototype           `json:"prices"`
	CreatedAt   *time.Time                 `json:"createdAt"`
	UpdatedAt   *time.Time                 `json:"updatedAt"`
}
