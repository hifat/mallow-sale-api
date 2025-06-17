package recipe

import (
	"time"

	"github.com/hifat/mallow-sale-api/internal/inventory"
	"github.com/hifat/mallow-sale-api/internal/usageUnit"
)

type (
	RecipeInventoryReq struct {
		UsageQuantity float64 `validate:"required" json:"usageQuantity"` // ปริมาณที่ใช้
		InventoryID   string  `validate:"required" json:"inventoryID"`   // ID สิ่งของ
		Remark        string  `json:"remark"`                            // หมายเหตุ

		UsageUnitCode string                   `validate:"required" json:"usageUnitCode"` // หน่วยใช้
		UsageUnit     usageUnit.UsageUnitEmbed `json:"-"`
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
		ID            string  `fake:"{uuid}" json:"id"`
		UsageQuantity float64 `fake:"{float64}" json:"usageQuantity"`
		Remark        string  `fake:"{sentence}" json:"remark"`

		UsageUnit *usageUnit.UsageUnitProtoType `json:"usageUnit"`

		InventoryID string                        `fake:"{uuid}" json:"-"`
		Inventory   *inventory.InventoryPrototype `json:"inventory"`
	}
)

type RecipeRes struct {
	ID        string     `fake:"{uuid}" json:"id"`
	Name      string     `fake:"{name}" json:"name"`
	CreatedAt *time.Time `fake:"{date}" json:"createdAt"`
	UpdatedAt *time.Time `fake:"{date}" json:"updatedAt"`

	Ingredients []RecipeInventoryRes `json:"ingredients"`
}

// Will return InventoryID in Ingredients
func (r *RecipeRes) GetInventoryIDs() []string {
	inventoryIDs := make([]string, 0, len(r.Ingredients))
	if r.Ingredients != nil {

		for _, v := range r.Ingredients {
			inventoryIDs = append(inventoryIDs, v.InventoryID)
		}
	}

	return inventoryIDs
}
