package recipeModule

import (
	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IngredientEntity struct {
	ID          primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	InventoryID primitive.ObjectID     `bson:"inventory_id" json:"inventoryID"`
	Inventory   inventoryModule.Entity `bson:"inventory" json:"inventory"`
	Quantity    float32                `bson:"quantity" json:"quantity"`
	Unit        usageUnitModule.Entity `bson:"unit" json:"unit"`
}

type Entity struct {
	utilsModule.Base `bson:"inline"`

	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name           string             `bson:"name" json:"name"`
	CostPercentage float32            `bson:"cost_percentage" json:"costPercentage"`
	Price          float32            `bson:"price" json:"price"`
	Ingredients    []IngredientEntity `bson:"ingredients" json:"ingredients"`
}
