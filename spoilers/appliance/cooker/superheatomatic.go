package cooker

import (
	"context"
	"fmt"
	"github.com/ProfessorMc/Recipe/spoilers/appliance"
	"github.com/ProfessorMc/Recipe/spoilers/dish"
	"sync"
	"time"
)

type SuperHeatOMatic struct {
	hasPower bool
	isOn bool
	currentTemp float32
	currentDish chan *dish.Dish
	completedDish chan *dish.Dish
	ctx context.Context
	cancel context.CancelFunc
}

func NewSuperHeatOMatic(capacity int) *SuperHeatOMatic {
	ctx, cancelFunc := context.WithCancel(context.Background())
	newHeatOMatic := &SuperHeatOMatic{
		ctx:ctx,
		cancel:cancelFunc,
		currentDish: make(chan *dish.Dish, capacity),
		completedDish: make(chan *dish.Dish, capacity),
	}
	newHeatOMatic.startDishHandlers(capacity)
	return newHeatOMatic
}

func (h *SuperHeatOMatic) SetPower(power bool) {
	h.hasPower = power
}

func (h *SuperHeatOMatic) TurnOn() error {
	if !h.hasPower {
		return appliance.BuildApplianceError("Not Plugged In", h)
	}
	h.isOn = true
	return nil
}

func (h *SuperHeatOMatic) TurnOff() error {
	h.isOn = false
	h.cancel()
	return nil
}

func (h SuperHeatOMatic) IsOn() bool {
	return h.isOn
}

func (SuperHeatOMatic) GetName() string {
	return "Super - Heat-O-Matic"
}

func (SuperHeatOMatic) GetBrand() string {
	return "Brandly"
}

func (h *SuperHeatOMatic) CookDish(d *dish.Dish) error {
	var wg sync.WaitGroup
	var err error
	go func() {
		err = h.CookDishAsync(d, &wg)
	}()
	wg.Wait()
	return err
}

func (h *SuperHeatOMatic) CookDishAsync(d *dish.Dish, wg *sync.WaitGroup) error {
	defer wg.Done()
	if !h.isOn {
		return appliance.BuildApplianceError("appliance isn't on", h)
	}
	// Push dish to current dish channel
	h.currentDish <- d
	// Wait for completed dish to finish
	<- h.completedDish
	return nil
}

func (h *SuperHeatOMatic) preheatOven(temp float32) {
	fmt.Printf("Preheating %s to %v degrees\n", h.GetName(), temp)
	<- time.After(5 * time.Second)
}

func (h *SuperHeatOMatic) startDishHandlers(capacity int) {
	fmt.Printf("[Appliance %s] Starting Dish Handlers.\n", h.GetName())

	for i:=0; i < capacity; i++ {
		started := make(chan struct{})
		go h.dishHandler(i, started)
		<-started
	}
}

func (h *SuperHeatOMatic) dishHandler(handlerNumber int, started chan struct{}) {
	fmt.Printf("[Appliance %s] Starting Handler %d.\n", h.GetName(), handlerNumber)
	close(started)
	for {
		select {
		case newDish := <- h.currentDish:
			fmt.Printf("[Appliance %s Handler %d] Handling Dish: %s\n", h.GetName(), handlerNumber, newDish.String())
			h.preheatOven(newDish.GetCookTemp())
			fmt.Printf("[Appliance %s Handler %d] Cooking Dish: %s\n", h.GetName(), handlerNumber, newDish.String())
			<- time.After(newDish.GetCookTime())
			newDish.SetTemperature(newDish.GetCookTemp())
			fmt.Printf("[Appliance %s Handler %d] Dish Complete: %s\n", h.GetName(), handlerNumber, newDish.String())
			h.completedDish <- newDish
		case <- h.ctx.Done():
			fmt.Printf("[Appliance %s Handler %d] Shutting Down Handler.\n", handlerNumber, h.GetName())
		}
	}
}



