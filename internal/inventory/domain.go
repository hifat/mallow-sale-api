package inventoryModule

import (
	"time"

	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
)

type Prototype struct {
	ID               string                    `json:"id"`
	Name             string                    `fake:"{name}" validate:"required" json:"name"`
	PurchasePrice    float32                   `fake:"{float32}" validate:"required" json:"purchase_price"`
	PurchaseQuantity float32                   `fake:"{float32}" validate:"required" json:"purchase_quantity"`
	PurchaseUnit     usageUnitModule.Prototype `json:"purchase_unit"`
	YieldPercentage  float32                   `fake:"{float32}" validate:"required" json:"yield_percentage"`
	Remark           string                    `fake:"{sentence}" json:"remark"`
	CreatedAt        *time.Time                `json:"created_at"`
	UpdatedAt        *time.Time                `json:"updated_at"`
}

type Request struct {
	Name             string                       `fake:"{name}" validate:"required" json:"name"`
	PurchasePrice    float32                      `fake:"{float32}" validate:"required" json:"purchase_price"`
	PurchaseQuantity float32                      `fake:"{float32}" validate:"required" json:"purchase_quantity"`
	PurchaseUnit     usageUnitModule.UsageUnitReq `json:"purchase_unit"`
	YieldPercentage  float32                      `fake:"{float32}" validate:"required" json:"yield_percentage"`
	Remark           string                       `fake:"{sentence}" json:"remark"`
}

type Response struct {
	Prototype
}
