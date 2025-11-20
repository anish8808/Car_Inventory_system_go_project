package main

import (
	"fmt"
	"time"
)

func printNumber() {
	for i := 0; i < 20; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("I m the other Goroutine")
}

func PrintWithSequenec(n int) {
	for i := 0; i <= n; i++ {
		fmt.Println(i)
		time.Sleep(2 * time.Second)
	}
}

func main() {

	//go printNumber() // this function will be kept in different goroutines ,
	// main wont wait for this go routines to compelete , for waiting we have to use time.sleep(10)

	go PrintWithSequenec(10)
	time.Sleep(10 * time.Second) // but only wait for 10 second so another go routines will work for 10 sec
	fmt.Println("I m the main Goroutine")
}
