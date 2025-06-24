package main

import "fmt"

func sender(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch) //close only have to use from sender
}

func main() {

	ch := make(chan int)

	go sender(ch)

	for val := range ch { // after hitting this line range on channel will stop once close(ch) used
		fmt.Println(val)
	}

	fmt.Println("Channel is closed ")

}
