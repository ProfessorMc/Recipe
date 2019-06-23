package cooker

import (
	"fmt"
	"github.com/ProfessorMc/Recipe/spoilers/appliance"
	"github.com/ProfessorMc/Recipe/spoilers/dish"
	"sync"
	"time"
)

type HeatOMatic struct {
	hasPower bool
	isOn bool
	currentTemp float32
	busy bool
	mtx sync.Mutex
}

func NewHeatOMatic() *HeatOMatic {
	newHeatOMatic := &HeatOMatic{
	}
	return newHeatOMatic
}

func (h *HeatOMatic) SetPower(power bool) {
	h.hasPower = power
}

func (h *HeatOMatic) TurnOn() error {
	if !h.hasPower {
		return appliance.BuildApplianceError("Not Plugged In", h)
	}
	h.isOn = true
	return nil
}

func (h *HeatOMatic) TurnOff() error {
	h.isOn = false
	return nil
}

func (h HeatOMatic) IsOn() bool {
	return h.isOn
}

func (HeatOMatic) GetName() string {
	return "Heat-O-Matic"
}

func (HeatOMatic) GetBrand() string {
	return "Brandly"
}

func (h *HeatOMatic) CookDish(d *dish.Dish) error {
	var wg sync.WaitGroup
	wg.Add(1)
	var err error
	go func() {
		err = h.CookDishAsync(d, &wg)
	}()
	wg.Wait()
	return err
}

func (h *HeatOMatic) CookDishAsync(d *dish.Dish, wg *sync.WaitGroup) error {
	defer wg.Done()
	if !h.isOn {
		return appliance.BuildApplianceError("appliance isn't on", h)
	}
	h.mtx.Lock()
	defer h.mtx.Unlock()
	fmt.Printf("[Appliance %s] Handling Dish: %s\n", h.GetName(), d.String())

	h.preheatOven(d.GetCookTemp())
	fmt.Printf("[Appliance %s] Cooking Dish: %s\n", h.GetName(), d.String())

	<- time.After(d.GetCookTime())
	fmt.Printf("[Appliance %s] Dish Complete: %s\n", h.GetName(), d.String())
	return nil
}

func (h *HeatOMatic) preheatOven(temp float32) {
	fmt.Printf("Preheating %s to %v degrees\n", h.GetName(), temp)
	<- time.After(5 * time.Second)
}





