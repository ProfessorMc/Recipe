package cooker

import (
	"github.com/ProfessorMc/Recipe/spoilers/dish"
	"sync"
)

type Cooker interface {
	CookDish(*dish.Dish)
	CookDishAsync(*dish.Dish, *sync.WaitGroup)
}
