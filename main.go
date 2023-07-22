package main

import (
	"fmt"
	"sync"
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

	example()
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

func example() {
	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			channel <- i
		}
		close(channel)
	}()

	for msg := range channel {
		fmt.Println(msg)
	}
}

func processByBatch() {
	c := make(chan int, 2)
	i := 0
	for {

		go func() {
			time.Sleep(time.Second)
			x := <-c
			println(x)
		}()
		fmt.Println("send ", i)
		c <- i
		i++
	}
}

func batchDoesntWork() {
	c := make(chan int)
	j := 0

	wg := sync.WaitGroup{}
	go func() {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(num *int) {
				defer wg.Done()
				fmt.Println("send ", *num)
				time.Sleep(time.Second)
				c <- *num
				*num++
			}(&j)
		}
		wg.Wait()
		close(c)
	}()

	for x := range c {
		fmt.Println(x)
	}
}

func waitGroupWithoutBufferChannel() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			time.Sleep(time.Second / 4)
			ch <- i
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	for c := range ch {
		fmt.Println(c)
	}

}

func sendBlocking() {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second / 2)
		fmt.Println(<-ch)
	}()
	go func() {
		time.Sleep(time.Second / 2)
		fmt.Println(<-ch)

	}()
	ch <- 1

	time.Sleep(time.Second * 3)
}
