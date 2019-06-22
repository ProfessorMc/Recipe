package main

import "fmt"
import a "./appliance"
import r "./recipe"
import d "./dish"
import t "./temperature"

// Lab 4.  Structs and Methods
// Requirements:
// 01 - As a lonely person living at the center of the Earth, I would like an application to the model my kitchen appliances
//
// Objective:
// 01 - Understand structs
// 02 - Understand understand methods
// 03 - Understand composition
// 04 - Run Basic Tests
//
// Steps:
//01 - Implement the dish name getter (TODO 1)
//02 - Implement the dish name setter (TODO 2)
//03 - Implement the temperature getter (TODO 3)
//04 - Implement the temperature setter (TODO 4)
//05 - Implement the dish to string method (TODO 5)
//06 - Implement the method to convert fahrenheit to celcuis (TODO 06)
//07 - Implement the method to convert fahrenheit to celcuis (TODO 07)
//08 - Implement the name getter (TODO 08)
//09 - Implement the brand getter (TODO 09)
//10 - Implement the method to turn on the chef-o-matic (TODO 10)
//11 - Implement the method to turn off the chef-o-matic (TODO 11)
//12 - Implement the method to determine if the chef-o-matic is on (TODO 12)
//13 - Implement the method for the chef-o-matic to make a dish from a recipe (TODO 13)
//14 - Implement the method to turn on the freeze-o-matic (TODO 14)
//15 - Implement the method to turn off the freeze-o-matic (TODO 15)
//16 - Implement the method to determine if the freeze-o-matic is on (TODO 16)
//17 - Implement the method for cooling a dish to a specific temperature (TODO 17)
//18 - Implement the method to turn on the heat-o-matic (TODO 18)
//19 - Implement the method to turn off the heat-o-matic (TODO 19)
//20 - Implement the method to determine if the heat-o-matic is on (TODO 20)
//21 - Implement the method for heating a dish to a specific temperature (TODO 21)

func main() {
	gcRecipe := r.New("Grilled Cheese Sandwich", 5.0, t.Fahrenheit(275), []string{"chedder cheese", "salsa", "bread"})

	chefOMatic := a.NewChefOMatic("Super Chef", "Awesome-0")
	heatOMatic := a.NewHeatOMatic("Star Flare", "Quasar")
	freezeOMatic := a.NewFreezeOMatic("Ice Ice Penguin", "Penguin Refridgurators")

	fmt.Println("Appiances found @ the center of the Earth")
	fmt.Println("-----------------------------------------")
	fmt.Println("Name: %s; Brand: %s", chefOMatic.Name(), chefOMatic.Brand())
	fmt.Println("Name: %s; Brand: %s", heatOMatic.Name(), heatOMatic.Brand())
	fmt.Println("Name: %s; Brand: %s", freezeOMatic.Name(), freezeOMatic.Brand())

	var gc *d.Dish

	if chefOMatic.On(); chefOMatic.IsOn() {
		gc = chefOMatic.Cook(gcRecipe)
		chefOMatic.Off()
	}

	fmt.Println("Cooked dish: %v", *gc)

	if heatOMatic.On(); heatOMatic.IsOn() {
		gc = heatOMatic.Heat(gc, t.Celcius(325))
		heatOMatic.Off()
	}

	fmt.Println("Heated dish: %v", *gc)

	if freezeOMatic.On(); freezeOMatic.IsOn() {
		gc = freezeOMatic.Cool(gc, t.Celcius(125))
		freezeOMatic.Off()
	}

	fmt.Println("Cooled dish: %v", *gc)
}
