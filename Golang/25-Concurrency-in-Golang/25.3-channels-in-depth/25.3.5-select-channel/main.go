package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)
	ch5 := make(chan string)

	go func() {
		time.Sleep(20 * time.Second)
		ch1 <- "Resonse of the channel 1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Resonse of the channel 2"
	}()
	go func() {
		time.Sleep(5 * time.Second)
		ch3 <- "Resonse of the channel 3"
	}()
	go func() {
		time.Sleep(4 * time.Second)
		ch4 <- "Resonse of the channel 4"
	}()
	go func() {
		time.Sleep(5 * time.Second) // also i can use default (for Non Blockinig)
		ch5 <- "Resonse of the channel 5"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case msg3 := <-ch3:
			fmt.Println(msg3)
		case msg4 := <-ch4:
			fmt.Println(msg4)
		case msg5 := <-ch5:
			fmt.Println(msg5)
		case <-time.After(1 * time.Second):
			fmt.Println("Time out as within the time no worker responded ")

		}
	}
}
