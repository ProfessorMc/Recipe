package dish

import t "../temperature"

type Dish struct {
	name string
	temp t.Fahrenheit
}

//TODO: 01 - Implement the dish name getter
func (d Dish) Name() string {
	name := ""

	return name
}

//TODO: 02 - Implement the dish name setter
func (d *Dish) SetName(name string) {

}

//TODO: 03 - Implement the temperature getter
func (d Dish) Temperature() t.Fahrenheit {
	var temp t.Fahrenheit

	return temp
}

//TODO: 04 - Implement the temperature setter
func (d *Dish) SetState(temp t.Fahrenheit) {

}

//TODO: 05 - Implement the dish to string method
func (d Dish) String() string {
	var s string

	return s
}
