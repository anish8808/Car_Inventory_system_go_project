package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string) // unbuffered channel , as here we havent mentioned the storage
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello Anish" //sender
	}()

	fmt.Println("Waiting for the sender ")
	msg := <-ch // reciver  ,  mian go routine will wait here until the sender sends something to this channel

	fmt.Println(msg)
}
