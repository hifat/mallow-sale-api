package inventory

import "time"

type (
	Inventory struct {
		ID               string     `json:"id"`
		Name             string     `json:"name"`             // วัตถุดิบ
		PurchasePrice    float64    `json:"purchasePrice"`    // ราคาซื้อ
		PurchaseQuantity float64    `json:"purchaseQuantity"` // ปริมาณซื้อ
		PurchaseUnit     UsageUnit  `json:"purchaseUnit"`     // หน่วยซื้อ
		YieldPercentage  float64    `json:"yieldPercentage"`  // Yield %
		UsageQuantity    float64    `json:"usageQuantity"`    // ปริมาณที่ใช้
		UsageUnit        UsageUnit  `json:"usageUnit"`        // หน่วยใช้
		Remark           string     `json:"remark"`           // หมายเหตุ
		CreatedAt        *time.Time `json:"createdAt"`
		UpdatedAt        *time.Time `json:"updatedAt"`
	}

	InventoryReq struct {
		Name             string    `validate:"required" json:"name"`             // วัตถุดิบ
		PurchasePrice    float64   `validate:"required" json:"purchasePrice"`    // ราคาซื้อ
		PurchaseQuantity float64   `validate:"required" json:"purchaseQuantity"` // ปริมาณซื้อ
		PurchaseUnit     UsageUnit `validate:"required" json:"purchaseUnit"`     // หน่วยซื้อ
		YieldPercentage  float64   `validate:"required" json:"yieldPercentage"`  // Yield %
		UsageQuantity    float64   `validate:"required" json:"usageQuantity"`    // ปริมาณที่ใช้
		UsageUnit        UsageUnit `validate:"required" json:"usageUnit"`        // หน่วยใช้
		Remark           string    `validate:"required,max=255" json:"remark"`   // หมายเหตุ
	}

	InventoryRes struct {
		Inventory
	}
)
