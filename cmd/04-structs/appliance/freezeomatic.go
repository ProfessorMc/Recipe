package appliance

import d "github.com/ProfessorMc/Recipe/cmd/04-structs/dish"
import t "github.com/ProfessorMc/Recipe/cmd/04-structs/temperature"

type FreezeOMatic struct {
	*Label
	activated bool
}

func NewFreezeOMatic(name, brand string) *FreezeOMatic {
	return &FreezeOMatic{Label: NewLabel(name, brand), activated: false}
}

//TODO: 14 - Implement the method to turn on the freeze-o-matic
func (fom *FreezeOMatic) On() {

}

//TODO: 15 - Implement the method to turn off the freeze-o-matic
func (fom *FreezeOMatic) Off() {

}

//TODO: 16 - Implement the method to determine if the freeze-o-matic is on
func (fom FreezeOMatic) IsOn() bool {
	var b bool

	return b
}

//TODO: 17 - Implement the method for cooling a dish to a specific temperature
//Bonus points of checking of the freeze-o-matic is on
func (fom FreezeOMatic) Cool(dish *d.Dish, newTemp t.Celcius) *d.Dish {
	var d d.Dish

	return &d
}
