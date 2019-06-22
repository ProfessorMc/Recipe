package appliance

import r "../recipe"
import d "../dish"

type ChefOMatic struct {
	*Label
	activated bool
}

func NewChefOMatic(name, brand string) *ChefOMatic {
	return &ChefOMatic{Label: NewLabel(name, brand), activated: false}
}

//TODO: 10 - Implement the method to turn on the chef-o-matic
func (com *ChefOMatic) On() {

}

//TODO: 11 - Implement the method to turn off the chef-o-matic
func (com *ChefOMatic) Off() {

}

//TODO: 12 - Implement the method to determine if the chef-o-matic is on
func (com ChefOMatic) IsOn() bool {
	var b bool

	return b
}

//TODO: 13 - Implement the method for the chef-o-matic to make a dish from a recipe
//Bonus points of checking of the chef-o-matic is on
func (com ChefOMatic) Cook(recipe *r.Recipe) *d.Dish {
	var d d.Dish

	return &d
}
