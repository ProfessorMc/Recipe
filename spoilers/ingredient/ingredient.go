package ingredient

type Ingredient struct {
	name   string
	sliced bool
	shouldBeSliced bool
}

func NewIngredient(name string, shouldBeSliced bool) *Ingredient {
	return &Ingredient{
		name: name,
		sliced: false,
		shouldBeSliced:shouldBeSliced,
	}
}

func (i Ingredient) Name() string {
	return i.name
}

func (i *Ingredient) Slice() {
	i.sliced = true
}

func (i Ingredient) Sliced() bool {
	return i.sliced
}

