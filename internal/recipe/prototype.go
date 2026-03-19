package recipeModule

import (
	"time"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
)

type IngredientPrototype struct {
	InventoryID string                     `json:"-"`
	Inventory   *inventoryModule.Prototype `json:"inventory"`

	Quantity float32                   `json:"quantity"`
	Unit     usageUnitModule.Prototype `json:"unit"`
}

type RecipeTypePrototype struct {
	Code EnumCodeRecipeType `json:"code"`
	Name string             `json:"name"`
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
	RecipeType      RecipeTypePrototype   `json:"recipeType"`
	No              int                   `json:"no"`
	Cost            float64               `json:"cost"`
	LinemanPrice    float64               `json:"linemanPrice"`
	GrabPrice       float64               `json:"grabPrice"`
}

func (p *Prototype) GetInventoryIDs() []string {
	inventoryIDs := make([]string, 0, len(p.Ingredients))
	for _, ingredient := range p.Ingredients {
		inventoryIDs = append(inventoryIDs, ingredient.InventoryID)
	}

	return inventoryIDs
}
