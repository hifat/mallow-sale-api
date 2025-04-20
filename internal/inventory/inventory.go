package inventory

import (
	"time"

	"github.com/hifat/cost-calculator-api/internal/usageUnit"
)

type (
	UsageUnitRes struct {
		ID   string `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
	}

	InventoryPrototype struct {
		ID               string     `json:"id"`
		Name             string     `json:"name"`
		PurchasePrice    float64    `json:"purchasePrice"`
		PurchaseQuantity float64    `json:"purchaseQuantity"`
		YieldPercentage  float64    `json:"yieldPercentage"`
		Remark           string     `json:"remark"`
		CreatedAt        *time.Time `json:"createdAt"`
		UpdatedAt        *time.Time `json:"updatedAt"`
	}

	InventoryReq struct {
		Name             string  `validate:"required" json:"name"`
		PurchasePrice    float64 `validate:"required" json:"purchasePrice"`
		YieldPercentage  float64 `validate:"required" json:"yieldPercentage"`
		Remark           string  `validate:"required,max=255" json:"remark"`
		PurchaseQuantity float64 `validate:"required" json:"purchaseQuantity"`

		PurchaseUnitCode string `validate:"required" json:"purchaseUnitCode"`
		PurchaseUnit     usageUnit.UsageUnitEmbed
	}

	InventoryRes struct {
		InventoryPrototype
		PurchaseUnit *UsageUnitRes `json:"purchaseUnit"`
	}
)
