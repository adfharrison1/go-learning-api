package main

import (
	"fmt"
	"time"
)

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("hello, ", phrase)

	doneChan <- true
}

func greet(phrase string, doneChan chan bool) {
	fmt.Println("hello, ", phrase)

	doneChan <- true
}

func main() {
	dones := make([]chan bool, 2)
	for i := range dones {
		dones[i] = make(chan bool)
	}

	go slowGreet("you slow bastard", dones[0])
	go greet("you quick fox", dones[1])

	for _, done := range dones {
		<-done
	}
}
