//Abstraction maens exposing only the neccessary parts of an object while hiding details

package main

import "fmt"

type shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func printArea(s shape) {
	fmt.Println(s.Area())
}

func main() {
	c := Circle{radius: 1.2}
	printArea(c) // here user are not aware how this printArea func is implemented
}
