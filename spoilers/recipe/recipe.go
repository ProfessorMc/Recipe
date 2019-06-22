package recipe

import "time"

type Recipe struct {
	name        string
	cookTime    time.Duration //Bonus points if you change this to an actual time type
	cookTemp    float32
	ingredients []string
}

func NewRecipe(name string, cookTime time.Duration, temp float32, ingredients ...string) *Recipe {
	return &Recipe{
		name:name,
		cookTime:cookTime,
		cookTemp:temp,
		ingredients:ingredients,
	}
}

func (r *Recipe) GetName() string {
	return r.name
}

func (r *Recipe) GetCookTime() time.Duration {
	return r.cookTime
}

func (r *Recipe) GetCookTemp() float32 {
	return r.cookTemp
}

func (r *Recipe) GetIngredients() []string {
	return r.ingredients
}

