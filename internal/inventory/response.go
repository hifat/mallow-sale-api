package inventoryModule

import (
	"time"

	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
)

type Prototype struct {
	ID               string                    `fake:"{uuid}" json:"id"`
	Name             string                    `fake:"{name}" json:"name"`
	PurchasePrice    float64                   `fake:"{float64}" json:"purchasePrice"`
	PurchaseQuantity float64                   `fake:"{float64}" json:"purchaseQuantity"`
	PurchaseUnit     usageUnitModule.Prototype `json:"purchaseUnit"`
	YieldPercentage  float32                   `fake:"{float32}" json:"yieldPercentage"`
	Remark           string                    `fake:"{sentence}" json:"remark"`
	CreatedAt        *time.Time                `json:"createdAt"`
	UpdatedAt        *time.Time                `json:"updatedAt"`
}

type Response struct {
	Prototype
}
