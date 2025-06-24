package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	go func() {

		ch <- 1
		fmt.Println("sending message 1 on channel...")
		ch <- 2
		fmt.Println("sending message 2 on channel...")
		ch <- 3
		fmt.Println("sending message 3 on channel...")
	}()

	fmt.Println("Waiting the sender")
	time.Sleep(5 * time.Second)
	i1 := <-ch
	fmt.Println("recived messgae :", i1)
	i2 := <-ch
	fmt.Println("recived messgae :", i2)
	i3 := <-ch
	fmt.Println("recived messgae :", i3)
	fmt.Println(i1, i2, i3)
}
