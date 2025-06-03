package recipeHandler

type Handler struct {
	RecipeRest *recipeRest
}

func New(RecipeRest *recipeRest) Handler {
	return Handler{RecipeRest}
}
