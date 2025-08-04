package inventoryModule

import (
	"time"

	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
)

type Prototype struct {
	ID               string                    `json:"id"`
	Name             string                    `fake:"{name}" json:"name"`
	PurchasePrice    float32                   `fake:"{float32}" json:"purchasePrice"`
	PurchaseQuantity float32                   `fake:"{float32}" json:"purchaseQuantity"`
	PurchaseUnit     usageUnitModule.Prototype `json:"purchaseUnit"`
	YieldPercentage  float32                   `fake:"{float32}" json:"yieldPercentage"`
	Remark           string                    `fake:"{sentence}" json:"remark"`
	CreatedAt        *time.Time                `json:"createdAt"`
	UpdatedAt        *time.Time                `json:"updatedAt"`
}

type Request struct {
	Name            string                       `fake:"{name}" validate:"required" json:"name"`
	PurchaseUnit    usageUnitModule.UsageUnitReq `json:"purchaseUnit"`
	YieldPercentage float32                      `fake:"{float32}" validate:"required" json:"yieldPercentage"`
	Remark          string                       `fake:"{sentence}" json:"remark"`
}

type Response struct {
	Prototype
}
