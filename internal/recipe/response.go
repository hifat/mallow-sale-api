package recipeModule

type Response struct {
	Prototype
}

type RecipeTypeResponse struct {
	Code EnumCodeRecipeType `json:"code"`
	Name string             `json:"name"`
}
