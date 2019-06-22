package robot

import i "../ingredient"

type JulienneRobot struct {
	name string
}

func NewJulienneRobot(name string) *JulienneRobot {
	return &JulienneRobot{name: name}
}

func (jr JulienneRobot) Name() string {
	return jr.name
}

//TODO 04 - Implement the Slice method for the Chopping Robot
func (jr JulienneRobot) Slice(ingredient *i.Ingredient) string {
	action := ""

	return action
}
