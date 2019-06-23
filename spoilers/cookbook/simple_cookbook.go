package cookbook

import (
	"errors"
	"fmt"
	"github.com/ProfessorMc/Recipe/spoilers/recipe"
)

type CookBook struct {
	name string
	recipes map[string]*recipe.Recipe
}

func NewCookBook(name string) *CookBook {
	return &CookBook{
		name:name,
		recipes:make(map[string]*recipe.Recipe),
	}
}

func (c *CookBook) AddRecipe(recipe *recipe.Recipe) error {
	if _, ok := c.recipes[recipe.GetName()]; ok {
		return errors.New("recipe already exists")
	}

	c.recipes[recipe.GetName()] = recipe
	return nil
}

func (c *CookBook) DeleteRecipe(name string) error {
	delete(c.recipes, name)
	return nil
}

func (c *CookBook) GetRecipe(name string) (*recipe.Recipe, error) {
	if _, ok := c.recipes[name]; !ok {
		return nil, errors.New("recipe already exists")
	}
	return c.recipes[name], nil
}

func (c *CookBook) PrintAll() {
	fmt.Printf("-- Cookbook: %s --\n", c.name)
	fmt.Println("Recipes:")
	for key, val := range c.recipes {
		fmt.Println("---------------------")
		fmt.Printf("--> NAME: %s\n", key)
		fmt.Printf("--> DIRECTIONS: Bake for %v @ %v\n", val.GetCookTime(), val.GetCookTemp())
		fmt.Printf("--> INGREDIENTS: %v\n", val.GetIngredients())
		fmt.Println("---------------------")

	}
}