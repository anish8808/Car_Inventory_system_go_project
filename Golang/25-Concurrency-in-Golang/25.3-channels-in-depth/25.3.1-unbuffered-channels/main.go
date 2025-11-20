package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string) // unbuffered channel , as here we havent mentioned the storage
	ch1 := make(chan string)

	go func() {
		for i := 0; i <= 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello Anish" //sender
	}()

	fmt.Println("Waiting for the sender ")
	msg := <-ch // reciver  ,  mian go routine will wait here until the sender sends something to this channel
	for i := range ch1 {
		fmt.Println(i)
	}
	fmt.Println(msg)
}
