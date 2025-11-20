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
The Singleton Design Pattern ensures that a class has only one instance and
provides a global point of access to it.
This pattern is useful when you need exactly one object (like a config or logger) shared across your application,
 and you want to prevent multiple instances from being created, especially in concurrent code.
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
