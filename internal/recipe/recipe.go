package recipeDomain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RecipeRepository interface {
	Create(recipe *Recipe) error
	GetByID(id primitive.ObjectID) (*Recipe, error)
	List() ([]*Recipe, error)
	Update(recipe *Recipe) error
	Delete(id primitive.ObjectID) error
}
