// Polymorphism means "many forms " in Go we can achived it by interface
package main

import "fmt"

type Animal interface {
	speak() string
}

type Dog struct {
}

func (d Dog) speak() string {
	return "bhow bhow"
}

type Cat struct {
}

func (c Cat) speak() string {
	return "meow meow"
}

func makeitSpeak(a Animal) {
	fmt.Println(a.speak())
}

func main() {
	makeitSpeak(Dog{})
	makeitSpeak(Cat{})
}
