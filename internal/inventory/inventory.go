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
		Name             string     `json:"name"`             // วัตถุดิบ
		PurchasePrice    float64    `json:"purchasePrice"`    // ราคาซื้อ
		PurchaseQuantity float64    `json:"purchaseQuantity"` // ปริมาณซื้อ
		YieldPercentage  float64    `json:"yieldPercentage"`  // Yield %
		UsageQuantity    float64    `json:"usageQuantity"`    // ปริมาณที่ใช้
		Remark           string     `json:"remark"`           // หมายเหตุ
		CreatedAt        *time.Time `json:"createdAt"`
		UpdatedAt        *time.Time `json:"updatedAt"`
	}

	InventoryReq struct {
		Name            string  `validate:"required" json:"name"`            // วัตถุดิบ
		PurchasePrice   float64 `validate:"required" json:"purchasePrice"`   // ราคาซื้อ
		YieldPercentage float64 `validate:"required" json:"yieldPercentage"` // Yield %
		Remark          string  `validate:"required,max=255" json:"remark"`  // หมายเหตุ

		PurchaseQuantity float64       `validate:"required" json:"purchaseQuantity"` // ปริมาณซื้อ
		PurchaseUnit     *UsageUnitRes `validate:"required" json:"purchaseUnit"`     // หน่วยซื้อ

		UsageQuantity float64       `validate:"required" json:"usageQuantity"` // ปริมาณที่ใช้
		UsageUnit     *UsageUnitRes `validate:"required" json:"usageUnit"`     // หน่วยใช้
	}

	InventoryRes struct {
		InventoryPrototype
		PurchaseUnit *UsageUnitRes `json:"purchaseUnit"` // หน่วยซื้อ
		UsageUnit    *UsageUnitRes `json:"usageUnit"`    // หน่วยใช้
	}
)
