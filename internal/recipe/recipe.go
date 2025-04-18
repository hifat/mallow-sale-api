package recipe

import (
	"time"

	"github.com/hifat/cost-calculator-api/internal/inventory"
	"github.com/hifat/cost-calculator-api/internal/usageUnit"
)

type (
	RecipeInventoryReq struct {
		UsageQuantity float64 `validate:"required" json:"usageQuantity"` // ปริมาณที่ใช้
		InventoryID   string  `validate:"required" json:"inventoryID"`   // ID สิ่งของ
		UsageUnitCode string  `validate:"required" json:"usageUnit"`     // หน่วยใช้
		Remark        string  `json:"remark"`                            // หมายเหตุ
	}

	RecipeReq struct {
		Name        string               `validate:"required" json:"name"`
		Inventories []RecipeInventoryReq `json:"inventories"`
	}

	RecipeInventoryRes struct {
		ID            string     `json:"id"`
		UsageQuantity float64    `json:"usageQuantity"`
		Remark        string     `json:"remark"`
		CreatedAt     *time.Time `json:"createdAt"`
		UpdatedAt     *time.Time `json:"updatedAt"`

		UsageUnit usageUnit.UsageUnitProtoType `json:"usageUnit"`
		Inventory inventory.InventoryPrototype `json:"inventory"`
	}

	RecipeRes struct {
		ID        string     `json:"id"`
		Name      string     `json:"name"`
		CreatedAt *time.Time `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`

		Inventories []RecipeInventoryRes `json:"inventories"`
	}
)
