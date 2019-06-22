package appliance

type Label struct {
	name  string
	brand string
}

func NewLabel(name, brand string) *Label {
	return &Label{name: name, brand: brand}
}

//TODO: 08 - Implement the name getter
func (l Label) Name() string {
	name := ""
	return name
}

//TODO: 09 - Implement the brand getter
func (l Label) Brand() string {
	brand := ""
	return brand
}
