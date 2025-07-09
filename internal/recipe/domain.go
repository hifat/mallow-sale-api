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

type Request struct {
	Name           string              `validate:"required" json:"name"`
	CostPercentage float32             `validate:"required" json:"costPercentage"`
	Ingredients    []IngredientRequest `validate:"required,dive" json:"ingredients"`
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

type Prototype struct {
	ID             string                `json:"id"`
	Name           string                `json:"name"`
	CostPercentage float32               `json:"costPercentage"`
	Ingredients    []IngredientPrototype `json:"ingredients"`
	CreatedAt      *time.Time            `json:"createdAt"`
	UpdatedAt      *time.Time            `json:"updatedAt"`
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
