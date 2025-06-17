package recipe_test

import (
	"testing"

	"github.com/hifat/mallow-sale-api/internal/recipe"
	"github.com/stretchr/testify/assert"
)

func TestRecipe_GetInventoryIDs(t *testing.T) {
	t.Run("should return inventory IDs when ingredients exist", func(t *testing.T) {
		_recipeRes := recipe.RecipeRes{
			Ingredients: []recipe.RecipeInventoryRes{
				{
					InventoryID: "1",
				},
				{
					InventoryID: "2",
				},
			},
		}

		inventoryIDs := _recipeRes.GetInventoryIDs()

		assert.Equal(t, []string{"1", "2"}, inventoryIDs)
		assert.Len(t, inventoryIDs, 2)
	})

	t.Run("should return empty slice when ingredients is nil", func(t *testing.T) {
		_recipeRes := recipe.RecipeRes{
			Ingredients: nil,
		}

		inventoryIDs := _recipeRes.GetInventoryIDs()

		assert.NotNil(t, inventoryIDs)
		assert.Empty(t, inventoryIDs)
	})
}
