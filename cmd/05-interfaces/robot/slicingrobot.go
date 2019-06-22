package robot

import i "../ingredient"

type SlicingRobot struct {
	name string
}

func NewSlicingRobot(name string) *SlicingRobot {
	return &SlicingRobot{name: name}
}

func (sr SlicingRobot) Name() string {
	return sr.name
}

//TODO 01 - Implement the Slice method for the Slicing Robot
func (sr SlicingRobot) Slice(ingredient *i.Ingredient) string {
	action := ""

	return action
}
