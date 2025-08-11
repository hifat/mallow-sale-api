package recipeModule

import (
	"time"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
)

type IngredientRequest struct {
	InventoryID string                       `validate:"required" json:"inventoryID"`
	Quantity    float32                      `validate:"required" json:"quantity"`
	Unit        usageUnitModule.UsageUnitReq `validate:"required" json:"unit"`
}

type RecipeTypeRequest struct {
	Code string `validate:"required" json:"code"`
	Name string `json:"-"`
}

type Request struct {
	Name            string              `validate:"required" json:"name"`
	CostPercentage  float32             `validate:"required" json:"costPercentage"`
	OtherPercentage float32             `json:"otherPercentage"`
	Price           float32             `validate:"gte=0" json:"price"`
	Ingredients     []IngredientRequest `validate:"required,dive" json:"ingredients"`
	Type            RecipeTypeRequest   `validate:"required" json:"type"`
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

type IngredientPrototype struct {
	InventoryID string                     `json:"-"`
	Inventory   *inventoryModule.Prototype `json:"inventory"`

	Quantity float32                   `json:"quantity"`
	Unit     usageUnitModule.Prototype `json:"unit"`
}

type RecipeTypePrototype struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Prototype struct {
	ID              string                `json:"id"`
	Name            string                `json:"name"`
	CostPercentage  float32               `json:"costPercentage"`
	OtherPercentage float32               `json:"otherPercentage"`
	Price           float32               `json:"price"`
	Ingredients     []IngredientPrototype `json:"ingredients"`
	CreatedAt       *time.Time            `json:"createdAt"`
	UpdatedAt       *time.Time            `json:"updatedAt"`
	Type            RecipeTypePrototype   `json:"type"`
	No              int                   `json:"no"`
}

func (p *Prototype) GetInventoryIDs() []string {
	inventoryIDs := make([]string, 0, len(p.Ingredients))
	for _, ingredient := range p.Ingredients {
		inventoryIDs = append(inventoryIDs, ingredient.InventoryID)
	}

	return inventoryIDs
}

type Response struct {
	Prototype
}

type UpdateOrderNoRequest struct {
	ID      string `json:"id"`
	OrderNo int    `json:"orderNo"`
}

type TypeResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
