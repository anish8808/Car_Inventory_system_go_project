package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker No: %v Processing Started\n", i)
	//logic
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker No: %v Processing Completed\n", i)
	// or we can use dwg.Done()  at end of any go routines to decrement the counter
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i+1, &wg)
	}

	wg.Wait()
	fmt.Println("Main function Process")
}

/*

	but in this example go routines will be executre
*/
