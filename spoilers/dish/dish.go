package dish

import (
	"fmt"
	f "github.com/ProfessorMc/Recipe/spoilers/friend"
	r "github.com/ProfessorMc/Recipe/spoilers/recipe"
)

type Dish struct {
	friend f.Friend
	r.Recipe
	currentTemperature float32
}

func NewDish(friend f.Friend, recipe r.Recipe) *Dish{
	return &Dish{
		friend:friend,
		Recipe: recipe,
	}
}

func (d Dish) Name() string {
	return d.String()
}

func (d Dish) GetTemperature() float32 {
	return d.currentTemperature
}

func (d *Dish) SetTemperature(temperature float32) {
	d.currentTemperature = temperature
}

func (d Dish) String() string {
	return fmt.Sprintf("Dish: %s for %s", d.Recipe.GetName() ,d.friend.GetName())
}
