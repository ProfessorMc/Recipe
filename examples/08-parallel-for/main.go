package main

import "fmt"

type Recipe struct {
	Name     string
	CookTime float32
	CookTemp float32
}

func (r Recipe) String() string {
	return fmt.Sprintf("{Name: %s; CookTime: %g; CookTemp: %g}", r.Name, r.CookTime, r.CookTemp)
}

// Lab 08.  Embarassingly Parallel
// Requirements:
// 01 - As a lonely person living at the center of the Earth, I would like an application that mirrors my intellectual proccessing of recipes
//
// Objective:
// 01 - See goroutines and channels in action
// 02 - See an example of a buffered channel
// 03 - Understand how to do implement a parallel for loop
//
// Steps:
//01 - Study the code to understand what's going on

func parallelFor(recipes []Recipe, fun func(recipe Recipe) Recipe) {
	type done struct{}
	finished := make(chan done, len(recipes))
	for index, value := range recipes {
		go func(i int, recipe Recipe) {
			recipes[i] = fun(recipe)
			finished <- done{}
		}(index, value)
	}

	for i := 0; i < len(recipes); i++ {
		<-finished
	}
}

func main() {
	//Make an array of 10 recipes and initialize them
	recipes := make([]Recipe, 10)
	for i := 0; i < len(recipes); i++ {
		recipes[i] = Recipe{Name: "Grilll Cheese", CookTime: 5.0, CookTemp: 275}
	}

	fmt.Printf("Before: %v\n\n", recipes)
	parallelFor(recipes, func(recipe Recipe) Recipe {
		recipe.CookTime *= 2
		recipe.CookTemp += 10
		return recipe
	})

	fmt.Printf("After: %v\n", recipes)
}
