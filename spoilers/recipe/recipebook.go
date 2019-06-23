package recipe

type RecipeBook interface {
	AddRecipe(recipe *Recipe) error
	DeleteRecipe(name string) error
	GetRecipe(name string) (*Recipe, error)
	PrintAll()
}
