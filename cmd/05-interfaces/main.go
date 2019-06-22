package main

import "fmt"
import i "github.com/ProfessorMc/Recipe/cmd/05-interfaces/ingredient"
import r "github.com/ProfessorMc/Recipe/cmd/05-interfaces/robot"

// Lab 5.  Interface, the Empty Interface{}, Type Assertions
// Requirements:
// 01 - As a lonely person living at the center of the Earth, I would like an application to the model my slicing kitchen robots
//
// Objective:
// 01 - Understand interfaces
// 02 - Understand the empty interface
// 03 - Understand type assertions
// 04 - Run Basic Tests
//
// Steps:
//01 - Implement the Slice method for the Slicing Robot (TODO 01)
//02 - Implement the Slice method for the Dicing Robot (TODO 02)
//03 - Implement the Slice method for the Chopping Robot (TODO 03)
//04 - Implement the Slice method for the Julienne Robot (TODO 04)

func performAction(robotName string, slicer r.Slicer, ingredient *i.Ingredient) {
	action := slicer.Slice(ingredient)
	fmt.Printf("Robot Name: %s; Action: %s; Ingredient Name: %s; Ingredent Sliced: %t\n",
		robotName, action, ingredient.Name(), ingredient.Sliced())
}

func listRobotNames(robots []interface{}) {
	for robot := range robots {
		printRobotName(robot)
	}
}

func printRobotName(robot interface{}) {
	switch v := robot.(type) {
	case r.SlicingRobot:
		fmt.Println(v.Name())
	case r.DicingRobot:
		fmt.Println(v.Name())
	case r.ChoppingRobot:
		fmt.Println(v.Name())
	case r.JulienneRobot:
		fmt.Println(v.Name())
	default:
		fmt.Println("Name Unknown")
	}
}

func main() {
	//Create Ingredients
	pepper := i.New("ghost pepper")
	carrot := i.New("carrot")
	onion := i.New("onion")
	celery := i.New("celery")

	//Create Robots
	slicer := r.NewSlicingRobot("Slicer")
	dicer := r.NewDicingRobot("Dicer")
	chopper := r.NewChoppingRobot("Chopper")
	julienner := r.NewJulienneRobot("Julienne")

	robots := []interface{}{slicer, dicer, chopper, julienner}

	//Perform actions
	performAction(slicer.Name(), slicer, pepper)
	performAction(dicer.Name(), dicer, onion)
	performAction(chopper.Name(), chopper, celery)
	performAction(julienner.Name(), julienner, carrot)

	//List Robot Names
	listRobotNames(robots)
}
