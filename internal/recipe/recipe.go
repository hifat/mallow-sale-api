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
		Remark        string  `json:"remark"`                            // หมายเหตุ

		UsageUnitCode string `validate:"required" json:"usageUnit"` // หน่วยใช้
		UsageUnit     usageUnit.UsageUnitEmbed
	}

	RecipeReq struct {
		Name        string               `validate:"required" json:"name"`
		Ingredients []RecipeInventoryReq `validate:"gt=0" json:"ingredients"`
	}

	UpdateRecipeInventoryReq struct {
		ID string `json:"id"`
		RecipeInventoryReq
	}

	UpdateRecipeReq struct {
		Name        string                     `validate:"required" json:"name"`
		Ingredients []UpdateRecipeInventoryReq `json:"ingredients"`
	}

	RecipeInventoryRes struct {
		ID            string  `json:"id"`
		UsageQuantity float64 `json:"usageQuantity"`
		Remark        string  `json:"remark"`

		UsageUnit *usageUnit.UsageUnitProtoType `json:"usageUnit"`

		InventoryID string                        `json:"inventoryID"`
		Inventory   *inventory.InventoryPrototype `json:"inventory"`
	}

	RecipeRes struct {
		ID        string     `json:"id"`
		Name      string     `json:"name"`
		CreatedAt *time.Time `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`

		Ingredients []RecipeInventoryRes `json:"ingredients"`
	}
)
