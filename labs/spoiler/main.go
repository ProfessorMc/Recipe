package main

import (
	"fmt"
)

func main() {
	doLab(1, lab1)
	doLab(2, lab2)
	doLab(3, lab3)
	doLab(4, lab4)
	doLab(5, lab5)
	doLab(6, lab6)
	doLab(7, lab7)
}

func doLab(labNum int, lab func()) {
	fmt.Printf("*********************** START LAB %d ***********************\n", labNum)
	lab()
	fmt.Printf("*********************** END LAB %d   ***********************\n\n\n", labNum)
}
