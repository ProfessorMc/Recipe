package recipe

import t "github.com/ProfessorMc/Recipe/cmd/04-structs/temperature"

type Recipe struct {
	name        string
	cookTime    float32 //Bonus points if you change this to an actual time type
	cookTemp    t.Fahrenheit
	ingredients []string
}

func New(name string, cookTime float32, cookTemp t.Fahrenheit, ingredients []string) *Recipe {
	return &Recipe{name: name, cookTime: cookTime, cookTemp: cookTemp, ingredients: ingredients}
}
