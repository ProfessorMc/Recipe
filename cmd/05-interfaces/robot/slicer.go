package robot

import i "github.com/ProfessorMc/Recipe/cmd/05-interfaces/ingredient"

type Slicer interface {
	Slice(ingredient *i.Ingredient) string
}
