package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string)
	go func() { messages <- "ping" }()
	go func() { messages <- "ping1" }()
	go func() { messages <- "ping2" }()
	go func() { messages <- "ping3" }()
	msg := <-messages
	fmt.Println(msg)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	done := make(chan bool, 1)
	go worker(done)
	//Block until we receive a notification from the worker on the channel.
	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

// The pong function accepts one channel for receives (pings) and a second for sends (pongs).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
