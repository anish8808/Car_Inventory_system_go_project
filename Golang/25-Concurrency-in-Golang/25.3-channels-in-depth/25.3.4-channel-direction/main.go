package main

import (
	"fmt"
	"time"
)

func send(ch chan<- int) {
	fmt.Printf("Sending on the channel:")
	time.Sleep(5 * time.Second)
	ch <- 30
}

func reciver(ch <-chan int) {
	fmt.Println("Waiting for the message...")
	msg := <-ch
	fmt.Println("Recived the message :", msg)
}

func main() {

	ch := make(chan int)
	go send(ch)

	reciver(ch)

}
