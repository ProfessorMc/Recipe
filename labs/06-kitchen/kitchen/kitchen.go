package kitchen

import (
	"fmt"
	"github.com/ProfessorMc/Recipe/spoilers/appliance"
)

type Kitchen struct {
	name string
	appliances []appliance.Appliance
}

func NewKitchen(name string, appliances ...appliance.Appliance) *Kitchen{
	return &Kitchen{
		name:name,
		appliances:appliances,
	}
}

func (k *Kitchen) PrintAppliances() {
	fmt.Printf("---- In %s we have: ---\n", k.name)
	for _, app := range k.appliances {
		fmt.Printf("----> %s (Brand: %s) \n", app.GetName(), app.GetBrand())
	}
}