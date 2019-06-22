package robot

import i "../ingredient"

type Slicer interface {
	Slice(ingredient *i.Ingredient) string
}
