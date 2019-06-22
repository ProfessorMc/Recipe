package recipe

import (
	"errors"
	"fmt"
)

type CookBook struct {
	name string
	recipes map[string]*Recipe
}

func NewCookBook(name string) *CookBook {
	return &CookBook{
		name:name,
		recipes:make(map[string]*Recipe),
	}
}

func (c *CookBook) AddRecipe(recipe *Recipe) error {
	if _, ok := c.recipes[recipe.name]; ok {
		return errors.New("recipe already exists")
	}

	c.recipes[recipe.name] = recipe
	return nil
}

func (c *CookBook) DeleteRecipe(name string) error {
	delete(c.recipes, name)
	return nil
}

func (c *CookBook) GetRecipe(name string) (*Recipe, error) {
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
		fmt.Printf("--> DIRECTIONS: Bake for %v @ %v\n", val.cookTime, val.cookTemp)
		fmt.Printf("--> INGREDIENTS: %v\n", val.ingredients)
		fmt.Println("---------------------")

	}
}