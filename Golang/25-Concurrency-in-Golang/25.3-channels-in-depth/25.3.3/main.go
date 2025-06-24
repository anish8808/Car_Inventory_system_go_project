package main

import (
	"fmt"
	"time"
)

func worker(i int, ch chan int) {
	fmt.Println("sending messgae to the channel", i)
	ch <- i
}

func main() {
	ch := make(chan int)

	for i := 0; i < 3; i++ {
		go worker(i, ch)
	}

	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Println("Recived message on channel: ", val)
	}
}
