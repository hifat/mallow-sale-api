package recipeModule

import (
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
)

type QueryReq struct {
	utilsModule.QueryReq

	RecipeTypeCode string `form:"recipeTypeCode"`
}

type IngredientRequest struct {
	InventoryID string                       `validate:"required" json:"inventoryID"`
	Quantity    float32                      `validate:"required" json:"quantity"`
	Unit        usageUnitModule.UsageUnitReq `validate:"required" json:"unit"`
}

type RecipeTypeRequest struct {
	Code EnumCodeRecipeType `validate:"required" json:"code"`
	Name string             `json:"-"`
}

type Request struct {
	Name            string              `validate:"required" json:"name"`
	CostPercentage  float32             `validate:"required" json:"costPercentage"`
	OtherPercentage float32             `json:"otherPercentage"`
	Price           float32             `validate:"gte=0" json:"price"`
	Ingredients     []IngredientRequest `validate:"required,dive" json:"ingredients"`
	RecipeType      RecipeTypeRequest   `validate:"required" json:"recipeType"`
	No              int                 `json:"no"`
}

func (r *Request) GetUsageUnitCodes() []string {
	usageUnitCodes := make([]string, 0, len(r.Ingredients))
	for _, ingredient := range r.Ingredients {
		usageUnitCodes = append(usageUnitCodes, ingredient.Unit.Code)
	}

	return usageUnitCodes
}

func (r *Request) GetInventoryIDs() []string {
	inventoryIDs := make([]string, 0, len(r.Ingredients))
	for _, ingredient := range r.Ingredients {
		inventoryIDs = append(inventoryIDs, ingredient.InventoryID)
	}

	return inventoryIDs
}

type UpdateOrderNoRequest struct {
	ID      string `json:"id"`
	OrderNo int    `json:"orderNo"`
}
