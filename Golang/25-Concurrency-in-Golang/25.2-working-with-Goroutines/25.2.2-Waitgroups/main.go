package main

import (
	"fmt"
	"sync"
	"time"
)

func workers(i, delay int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Wokerr no: %v processing started\n", i)

	time.Sleep(time.Duration(delay) * time.Second)

	fmt.Printf("worker no : %v processing Completed\n", i)
}
func main() {

	var wg sync.WaitGroup
	wg.Add(5) // counter 5
	go workers(1, 2, &wg)
	go workers(2, 3, &wg)
	go workers(3, 1, &wg)
	go workers(4, 2, &wg)
	go workers(5, 3, &wg)

	wg.Wait()
}
