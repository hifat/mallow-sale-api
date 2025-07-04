package recipeIngredientDomain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RecipeIngredient struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	RecipeID    primitive.ObjectID `bson:"recipe_id" json:"recipe_id"`
	InventoryID primitive.ObjectID `bson:"inventory_id" json:"inventory_id"`
	Amount      int                `bson:"amount" json:"amount"`
}

type RecipeIngredientRepository interface {
	Create(ri *RecipeIngredient) error
	GetByID(id primitive.ObjectID) (*RecipeIngredient, error)
	ListByRecipe(recipeID primitive.ObjectID) ([]*RecipeIngredient, error)
	Update(ri *RecipeIngredient) error
	Delete(id primitive.ObjectID) error
}
