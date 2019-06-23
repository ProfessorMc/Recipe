package main

import "github.com/ProfessorMc/Recipe/spoilers/dish"

var dishes []*dish.Dish

func lab5(){
	dishes = make([]*dish.Dish,0)
	for _, friend := range friends {
		cakeRecipe, err := cookbook.GetRecipe("cake")
		if err != nil {
			panic(err)
		}
		order := dish.NewDish(*friend, *cakeRecipe)
		dishes = append(dishes, order)
	}
}