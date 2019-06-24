package main

import (
	"fmt"
	"sync"
)

var messagesReceived = 0

func main() {
	msgChan := make(chan string, 10)
	doneChan := make(chan struct{}, 0)
	var wg sync.WaitGroup

	go watchTotals(doneChan)
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("Message: %d", i)
		msgChan <- msg
		wg.Add(1)
		go startMessageReceiver(msgChan, doneChan, &wg, i)
	}

	//close(msgChan)
	//close(doneChan)
	wg.Wait()
}

func startMessageReceiver(msgChan chan string, doneChan chan struct{}, wg *sync.WaitGroup, receiverId int) {
	defer wg.Done()
	fmt.Println("Starting Receiver: ", receiverId)
	for {
		select {
		case msg := <-msgChan:
			fmt.Printf("[Receiver %d] Message: %s\n", receiverId, msg)
			messagesReceived++
		case <-doneChan:
			return
		}
	}
}

func watchTotals(doneChan chan struct{}) {
	for {
		if messagesReceived >= 10 {
			close(doneChan)
			return
		}
	}
}
