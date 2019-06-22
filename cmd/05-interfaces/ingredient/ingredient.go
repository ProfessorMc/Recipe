package ingredient

type Ingredient struct {
	name   string
	sliced bool
}

func New(name string) *Ingredient {
	return &Ingredient{name: name, sliced: false}
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
