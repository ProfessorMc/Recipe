package appliance

import d "github.com/ProfessorMc/Recipe/cmd/04-structs/dish"
import t "github.com/ProfessorMc/Recipe/cmd/04-structs/temperature"

type HeatOMatic struct {
	*Label
	activated bool
}

func NewHeatOMatic(name, brand string) *HeatOMatic {
	return &HeatOMatic{Label: NewLabel(name, brand), activated: false}
}

//TODO: 18 - Implement the method to turn on the heat-o-matic
func (hom *HeatOMatic) On() {

}

//TODO: 19 - Implement the method to turn off the heat-o-matic
func (hom *HeatOMatic) Off() {

}

//TODO: 20 - Implement the method to determine if the heat-o-matic is on
func (hom HeatOMatic) IsOn() bool {
	var b bool

	return b
}

//TODO: 21 - Implement the method for heating a dish to a specific temperature
func (hom HeatOMatic) Heat(dish *d.Dish, newTemp t.Celcius) *d.Dish {
	var d d.Dish

	return &d
}
