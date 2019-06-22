package main

import (
	"fmt"
)

// Lab 3.  Arrays, Slices, and Maps
// Requirements:
// 01 - As a lonely person, I would like an application to collect my favorate recipes
// 02 - As a lonely person, I would like the application to be able to index my recipes by name
//
// Objective:
// 01 - Understand array declarations
// 02 - Understand map declarations
// 03 - Understand slices and thier relationship to arrays
// 04 - Run Basic Tests
//
// Steps:
// 01 - Add 3 recipes to the recipes array (TODO 1)
// 02 - Modify the cookTime for the second recipe in the array (TODO 2)
// 03 - Print only the second recipe (TODO 3)
// 04 - Add a new recipe (TODO 4)
// 05 - Create a recipe index (Hint: map) (TODO 5)
// 06 - Remove a recipe from the recipe index (TODO 6)
// 07 - Run unit tests

type recipe struct {
	name        string
	cookTime    float32 //Bonus points if you change this to an actual time type
	cookTemp    float32
	ingredients []string
}

func updateCookTime(index int, recipes []recipe) {

}

func printRecipes(recipes []recipe) {
	fmt.Println("My Lonely Person Cookbook")
	fmt.Println("-------------------------")

	for i, v := range recipes {
		fmt.Printf("index: %d; %v\n", i, v)
	}
}

func addRecipe(recipes []recipe, newRecipe recipe) []recipe {
	return recipes
}

func createIndex(recipes []recipe) map[string]int {
	recipeIndex := make(map[string]int)

	return recipeIndex
}

func printRecipeIndex(recipeIndex map[string]int) {
	fmt.Println("My Lonely Person Cookbook Index")
	fmt.Println("-------------------------------")

	for i, v := range recipeIndex {
		fmt.Printf("index: %s; %d\n", i, v)
	}
}

func removeFromRecipeIndex(name string, recipeIndex map[string]int) map[string]int {

	return recipeIndex
}

func main() {
	// TODO: 01 - Add 3 recipes (bonus points for array literals)
	var recipes [3]recipe

	printRecipes(recipes[:])

	// TODO 02 - Modify the cook time for the second Recipe
	updateCookTime(2, recipes[:])

	//TODO 03 - Print only the second Recipe
	printRecipes(recipes[:])

	//TODO 04 - Add new recipe
	var newRecipe recipe
	expandedRecipes := addRecipe(recipes[:], newRecipe)

	printRecipes(expandedRecipes)

	//TODO 05 - Create a recipe index
	var recipeIndex = createIndex(expandedRecipes)

	printRecipeIndex(recipeIndex)

	//TODO 06 - Remove a recipe by name from the recipe index
	name := ""
	removeFromRecipeIndex(name, recipeIndex)

	printRecipeIndex(recipeIndex)
}
