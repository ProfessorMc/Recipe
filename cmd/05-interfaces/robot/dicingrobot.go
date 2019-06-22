package robot

import i "../ingredient"

type DicingRobot struct {
	name string
}

func NewDicingRobot(name string) *DicingRobot {
	return &DicingRobot{name: name}
}

func (dr DicingRobot) Name() string {
	return dr.name
}

//TODO 02 - Implement the Slice method for the Dicing Robot
func (dr DicingRobot) Slice(ingredient *i.Ingredient) string {
	action := ""

	return action
}
