package router

import (
	"github.com/hifat/cost-calculator-api/internal/recipe/recipeDI"
)

func (r *router) RecipeRouter() {
	handler := recipeDI.Init(r.db, r.logger, r.validator)

	recipe := r.route.Group("/api/recipes")
	recipe.Get("", handler.RecipeRest.Find)
	recipe.Get("/:recipeID", handler.RecipeRest.FindByID)
	recipe.Post("", handler.RecipeRest.Create)
	recipe.Put("/:recipeID", handler.RecipeRest.Update)
	recipe.Delete("/:recipeID", handler.RecipeRest.Delete)
}
