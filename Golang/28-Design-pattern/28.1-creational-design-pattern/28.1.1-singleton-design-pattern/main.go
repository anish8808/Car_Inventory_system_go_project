package main

import (
	"fmt"
	"sync"
)

//singleton
//factory
//adater
//stretegy
//observer
//decorater

/*

 */

type config struct {
	//here som value
}

var count = 0

var wg sync.WaitGroup
var mu sync.Mutex
var singeltonObject *config

func getConfigInstance() *config {
	defer wg.Done()
	if singeltonObject == nil {
		mu.Lock()
		defer mu.Unlock()
		if singeltonObject == nil {
			singeltonObject = &config{}
			count = count + 1
			fmt.Println("Created the single instance of the config", count)
		} else {
			fmt.Println("already Present the config Instance ")
		}
	} else {
		fmt.Println("already Present the config Instance ")
	}

	return singeltonObject
}

func main() {

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go getConfigInstance()
	}
	wg.Wait()
}
