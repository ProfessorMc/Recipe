package main

import (
	"fmt"
	"reflect"
)

// Lab 6.  A Time to Reflect
// Requirements:
// 01 - As a lonely person living at the center of the Earth, I would like an application that would allow me to relect on my RoboChef
//
// Objective:
// 01 - Understand how to use reflection to access values
// 02 - Understand how to use reflection to set values
// 03 - Understand struct tags
// 04 - Understand how to use reflection inorder to access struct tags
// 04 - Run Basic Tests
//
// Steps:
//01 - Use the Go lang documentation to describe the associated code (TODO 01)
//02 - Use the Go lang documentation to describe the associated code (TODO 02)
//03 - Use the Go lang documentation to describe the associated code (TODO 03)
//Reference:
//	https://blog.golang.org/laws-of-reflection

type RoboChef struct {
	Name         string   `RoboChefs Name`
	Make         string   `Company who made RoboChef`
	Model        string   `RoboChefs version model`
	Specialities []string `RoboChefs specialities`
}

func NewRoboChef(name, make, model string, specialities []string) *RoboChef {
	return &RoboChef{Name: name, Make: make, Model: model, Specialities: specialities}
}

func main() {
	roboChef := NewRoboChef("Ninja Chef", "Superior Robots", "Mark 12", []string{"slcing", "dicing", "stirfrying"})

	//Use reflection to access values (TODO 01)
	fmt.Println("Robot Type:", reflect.TypeOf(roboChef))
	fmt.Println("Robot Name:", reflect.ValueOf(roboChef.Name).String())
	fmt.Println("Robot Make:", reflect.ValueOf(roboChef.Make).String())
	fmt.Println("Robot Model:", reflect.ValueOf(roboChef.Model).String())

	specialities := reflect.ValueOf(roboChef.Specialities)

	fmt.Println("Robot Specialities:")
	for i := 0; i < specialities.Len(); i++ {
		speciality := specialities.Index(i).String()
		fmt.Println("\t", speciality)
	}

	//Update the model of RoboChef (TODO 02)
	value := reflect.ValueOf(roboChef)
	s := value.Elem() //Get the interface struct pointer
	model := s.FieldByName("Model")
	model.SetString("Mark 13")

	fmt.Println("Upgraded Robot Model:", reflect.ValueOf(roboChef.Model).String())

	//Get Struct Tags (TODO 03)
	t := reflect.TypeOf(roboChef).Elem()

	name, _ := t.FieldByName("Name")
	fmt.Println("Robot Name Tag:", name.Tag)

	make, _ := t.FieldByName("Make")
	fmt.Println("Robot Make Tag:", make.Tag)

	mod, _ := t.FieldByName("Model")
	fmt.Println("Robot Model Tag:", mod.Tag)

	specs, _ := t.FieldByName("Specialities")
	fmt.Println("Robot Specialities Tag:", specs.Tag)
}
