package main

import (
	"fmt"
	"sync"
	"time"
)

func workers(id int, tasks <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("Task %d is picked up by the worker %d\n", task, id)
		time.Sleep(1 * time.Second) //simulate the logic

		result <- task * task
	}
}
func main() {

	const worker = 3
	const jobs = 7
	tasks := make(chan int, jobs)
	result := make(chan int, jobs)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)
		go workers(i, tasks, result, &wg)
	}

	for i := 1; i < jobs; i++ {
		tasks <- i
	}

	close(tasks)

	wg.Wait()

	close(result)

	for r := range result {
		fmt.Println(r)
	}

}
