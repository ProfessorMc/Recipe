package robot

import i "../ingredient"

type ChoppingRobot struct {
	name string
}

func NewChoppingRobot(name string) *ChoppingRobot {
	return &ChoppingRobot{name: name}
}

func (cr ChoppingRobot) Name() string {
	return cr.name
}

//TODO 03 - Implement the Slice method for the Chopping Robot
func (cr ChoppingRobot) Slice(ingredient *i.Ingredient) string {
	action := ""

	return action
}
