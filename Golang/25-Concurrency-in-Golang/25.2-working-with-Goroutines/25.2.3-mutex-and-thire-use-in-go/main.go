package main

import (
	"fmt"
	"sync"
)

var (
	counter  int
	counterw int
	mu       sync.Mutex
	rw       sync.RWMutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	counter++
	mu.Unlock()
}

func readCounter() int {
	rw.RLock()
	defer rw.Unlock()

	return counter
}

func writeCounter() {
	rw.Lock()
	defer rw.Unlock()
	counter++
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()

	fmt.Println(counter)

	var wgrw sync.WaitGroup

	for i := 0; i < 10; i++ {
		wgrw.Add(1)
		go func() {
			wgrw.Done()
			fmt.Printf("Reading counter %d\n", readCounter())
		}()
	}

	for i := 0; i < 10; i++ {
		wgrw.Add(1)
		go func() {
			wgrw.Done()
			writeCounter()
		}()
	}

	wg.Wait()
}
