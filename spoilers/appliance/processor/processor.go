package processor

import "github.com/ProfessorMc/Recipe/spoilers/ingredient"

type Processor interface {
	ProcessIngredient(*ingredient.Ingredient)
}
