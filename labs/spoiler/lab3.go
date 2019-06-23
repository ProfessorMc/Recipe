package main

import (
	"fmt"
	"github.com/ProfessorMc/Recipe/spoilers/recipe"
	"time"
)


func lab3() {
	fmt.Println("-- Lab3 Recipes --")

	cake := recipe.NewRecipe("cake", time.Duration(3 * time.Second), 350, "cake mix", "eggs", "oil")
	ziti := recipe.NewRecipe("ziti", time.Duration(6 * time.Second), 300, "noodles", "water", "cheesy goodness")
	fish := recipe.NewRecipe("fish", time.Duration(6 * time.Second), 450, "fish, duh", "spice")

	fmt.Println("I can make: ", cake.GetName())
	fmt.Println("I can make: ", ziti.GetName())
	fmt.Println("I can make: ", fish.GetName())

}