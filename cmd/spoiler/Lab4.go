package main

import (
	"fmt"
	cb "github.com/ProfessorMc/Recipe/spoilers/cookbook"
	"github.com/ProfessorMc/Recipe/spoilers/recipe"
	"time"
)

var cookbook *cb.CookBook

func lab4() {
	fmt.Println("-- Here's what I can Make --")

	cake := recipe.NewRecipe("cake", time.Duration(3 * time.Second), 350, "cake mix", "eggs", "oil")
	ziti := recipe.NewRecipe("ziti", time.Duration(6 * time.Second), 300, "noodles", "water", "cheesy goodness")
	fish := recipe.NewRecipe("fish", time.Duration(6 * time.Second), 450, "fish, duh", "spice")

	cookbook = cb.NewCookBook("Muh Cookbook")
	_ = cookbook.AddRecipe(cake)
	_ = cookbook.AddRecipe(ziti)
	_ = cookbook.AddRecipe(fish)
	cookbook.PrintAll()

	fmt.Println("Crap, there's a fish allergy")
	_ = cookbook.DeleteRecipe("fish")

	fmt.Println("-- Here's what I can Make Tonight --")

	cookbook.PrintAll()

}