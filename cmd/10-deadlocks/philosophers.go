package main

import (
	"fmt"
	"time"
)

// Lab 10.  A Recipe for Dining Philosophers
// Requirements:
// 01 - As a lonely person living at the center of the Earth, I would like an application that would simulate how 5 philosophers dine in my impenetrable fortress
//
// Objective:
// 01 - See goroutines and channels in action
// 02 - See a potential deadlock scenario occur
// 03 - Understand why the deadlock occurs
// BONUS POINTS: Modify the code to remove the deadlock
//
// Steps:
//01 - Study the code to understand what's going on
//02 - Run the code a couple of times to observe the deadlock
//03 - Explain why go can sometimes detect the deadlock and sometimes not
//Reference:
//	https://en.wikipedia.org/wiki/Dining_philosophers_problem

type fork struct{}
type stop struct{}

func logger(log <-chan string, quit <-chan stop) string {
	for {
		select {
		case msg := <-log:
			fmt.Println(msg)
		case <-quit:
			break
		}
	}
}

func philosopher(name string, logger chan<- string, leftForkChan chan fork, rightForkChan chan fork, quit <-chan stop, eatTime time.Duration) {
	for {
		select {
		case leftFork := <-leftForkChan:
			logger <- fmt.Sprintf("Name: %s acquired left fork waiting on right fork", name)
			rightFork := <-rightForkChan
			logger <- fmt.Sprintf("Name: %s has both left and right forks, and is eating", name)
			time.Sleep(eatTime * time.Millisecond)
			logger <- fmt.Sprintf("Name: %s done eating releasing forks", name)
			rightForkChan <- rightFork
			leftForkChan <- leftFork

		case rightFork := <-rightForkChan:
			logger <- fmt.Sprintf("Name: %s acquired right fork waiting on left fork", name)
			leftFork := <-leftForkChan
			logger <- fmt.Sprintf("Name: %s has both left and right forks, and is eating", name)
			time.Sleep(eatTime * time.Millisecond)
			logger <- fmt.Sprintf("Name: %s done eating releasing forks", name)
			leftForkChan <- leftFork
			rightForkChan <- rightFork

		case <-quit:
			break
		}
	}
}

func timeout(timeout time.Duration, quit chan<- stop) {
	time.Sleep(timeout * time.Second)

	var kill stop
	quit <- kill
}

func main() {
	//Have the robot servants acquire the forks
	var forks [5]chan fork
	for i := range forks {
		forks[i] = make(chan fork)
	}

	quit := make(chan stop)
	log := make(chan string)
	go logger(log, quit)

	//Bring the philosophers to the table for a wonderful spagetti dinner which RoboChef has made
	names := [5]string{"Leonardo", "Michelangelo", "Donatello", "Raphael", "Sun Tzu"}
	for i := 0; i < 5; i++ {
		go philosopher(names[i], log, forks[i], forks[(i+1)%5], quit, 300)
	}

	go timeout(30, quit) //Start this before the deadlock situation is created

	//Provide the forks to the philosophers and let the fun begin
	for i := 0; i < 5; i++ {
		var newFork fork
		forks[i] <- newFork
	}

	//Watch and wait
	log <- "Main goroutine waiting on timeout"
	<-quit

	fmt.Println("Dinner is OVAH!!!")
}
