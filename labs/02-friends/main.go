package main

import (
	"fmt"
)

// Lab 2.  Friends - Variables, Functions, and Pointers
// Requirements:
// 01 - As a lonely person, I would like an application to greet my friends by name and age.
// 02 - As a caring child, I would like the app to remember my dad lies about his age and use his fake age (-5 years) in his greeting.
//
// Objective:
// 01 - Understand variable declarations
// 02 - Understand functions, pointers, types
// 03 - Run Basic Tests
// 04 - Import packages
// 05 - Use go doc
//
// Steps:
// Part 1. Implement Using Functions (Instructions in functions.go)
// Part 2. Implement Using Package (Instructions in packages.go)


func main() {
	fmt.Println("Running Lab 02 - Friends")
	UsingFunctions()
	UsingPackages()
}
