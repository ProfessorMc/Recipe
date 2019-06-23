package recipe

import "time"

type Recipe struct {

}

func NewRecipe(name string, cookTime time.Duration, temp float32, ingredients ...string) *Recipe {
	panic("implement NewRecipe")
}

func (r *Recipe) GetName() string {
	panic("implement NewRecipe")
}

func (r *Recipe) GetCookTime() time.Duration {
	panic("implement GetCookTime")
}

func (r *Recipe) GetCookTemp() float32 {
	panic("implement GetCookTemp")
}

func (r *Recipe) GetIngredients() []string {
	panic("implement GetIngredients")
}

