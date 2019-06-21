package main

import (
	"fmt"
)

// Lab 2.  Variables, Functions, and Pointers
// Requirements:
// 01 - As a lonely person, I would like an application to greet my friends by name and age.
// 02 - As a caring child, I would like the app to remeber my dad lies about his age and use his fake age (-5 years) in his greeting.
//
// Objective:
// 01 - Understand variable declarations
// 02 - Understand functions, pointers, types
// 03 - Run Basic Tests
//
// Steps:
// 01 - Add my best friend's name and age (TODO 1)
// 02 - Add my best mom's name and age (TODO 2)
// 03 - Add my dad's name and age (TODO 3)
// 04 - Refactor to support Dad's birthday (Hint: Pointer)
// 05 - Add a new property to friend to reflect if friend lies about age (Hint: bool, integer )
// 06 - Run unit tests

type friend struct {
	name string
	age  int
}

var myDad friend

func main() {
	// TODO: 01 - Declare my best friend's name and age
	myOnlyFriend := friend{}

	// TODO: 02 - Add my mom's name and age (21)
	var myMom friend

	// TODO: 03 - Add my dad's age and name

	greetFriend(myOnlyFriend)
	greetFriend(myMom)
	greetMyDad()
	fmt.Println("Oops, forgot dad's birthday...")
	celebrateBirthday(myDad)
	greetMyDad()
}

func greetFriend(f friend) {
	fmt.Printf("Hello %s, who is %d, years old.\n", f.name, f.age)
}

func celebrateBirthday(f friend) {
	f.age = addAge(f, 1)
}

func addAge(f friend, increment int) int {
	return f.age + increment
}

func greetMyDad() {
	greetFriend(myDad)
}
