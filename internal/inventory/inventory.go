package inventory

import "time"

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
		UsageQuantity    float64    `json:"usageQuantity"`
		Remark           string     `json:"remark"`
		CreatedAt        *time.Time `json:"createdAt"`
		UpdatedAt        *time.Time `json:"updatedAt"`
	}

	InventoryReq struct {
		Name            string  `validate:"required" json:"name"`
		PurchasePrice   float64 `validate:"required" json:"purchasePrice"`
		YieldPercentage float64 `validate:"required" json:"yieldPercentage"`
		Remark          string  `validate:"required,max=255" json:"remark"`

		PurchaseQuantity float64       `validate:"required" json:"purchaseQuantity"`
		PurchaseUnit     *UsageUnitRes `validate:"required" json:"purchaseUnit"`

		UsageQuantity float64       `validate:"required" json:"usageQuantity"`
		UsageUnit     *UsageUnitRes `validate:"required" json:"usageUnit"`
	}

	InventoryRes struct {
		InventoryPrototype
		PurchaseUnit *UsageUnitRes `json:"purchaseUnit"`
		UsageUnit    *UsageUnitRes `json:"usageUnit"`
	}
)
