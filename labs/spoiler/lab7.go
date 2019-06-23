
package main

import (
	"fmt"
	"github.com/ProfessorMc/Recipe/spoilers/appliance"
	"github.com/ProfessorMc/Recipe/spoilers/appliance/cooker"
	"github.com/ProfessorMc/Recipe/spoilers/kitchen"
	"sync"
	"time"
)


func lab7(){
	superOven := cooker.NewSuperHeatOMatic(5)
	muhKitchen := kitchen.NewKitchen("Muh Super Kitchen", superOven)
	muhKitchen.PrintAppliances()

	appliance.Plugin(superOven)
	err := superOven.TurnOn()
	if err != nil {
		panic(err)
	}

	startTime := time.Now()
	var wg sync.WaitGroup
	for dishIndex := range dishes {
		wg.Add(1)
		go superOven.CookDishAsync(dishes[dishIndex], &wg)
	}
	wg.Wait()
	cookTime := time.Since(startTime)
	fmt.Printf("Muh Super Kitchen Took %v to make %d recipes.\n", cookTime, len(dishes))
}


