package main

import (
	"fmt"
	"github.com/ProfessorMc/Recipe/spoilers/appliance"
	"github.com/ProfessorMc/Recipe/spoilers/appliance/cooker"
	"github.com/ProfessorMc/Recipe/spoilers/kitchen"
	"time"
	"log"
)

func lab6(){
	oven := cooker.NewHeatOMatic()
	muhKitchen := kitchen.NewKitchen("Muh Kitchen", oven)
	muhKitchen.PrintAppliances()
	startTime := time.Now()

	for dishIndex := range dishes {
		if err := oven.CookDish(dishes[dishIndex]); err != nil {
			log.Println(err)
		}
	}

	appliance.Plugin(oven)
	err := oven.TurnOn()
	if err != nil {
		panic(err)
	}

	for dishIndex := range dishes {
		if err := oven.CookDish(dishes[dishIndex]); err != nil {
			panic(err)
		}
	}

	cookTime := time.Since(startTime)
	fmt.Printf("Muh Kitchen Took %v to make %d recipes.\n", cookTime, len(dishes))
}
