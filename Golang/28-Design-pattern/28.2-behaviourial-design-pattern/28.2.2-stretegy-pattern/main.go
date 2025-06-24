package main

import "fmt"

/*
Strategy Design Pattern is one of the behavioural design pattern that is used when we have the multiple
algorithm for a specfic task and client decides the actual implementation to be used at run time , it also
none as a policy pattern
*/

type OperationStrategy interface {
	DoOperation(a, b int) int
}

type AddOperation struct{}

func (AddOperation) DoOperation(a, b int) int {
	return a + b
}

type SubtractOperation struct{}

func (SubtractOperation) DoOperation(a, b int) int {
	return a - b
}

type Calculator struct {
	stretgy OperationStrategy
}

func (c *Calculator) SetStrategy(s OperationStrategy) {
	c.stretgy = s
}

func (c *Calculator) calculate(a, b int) int {
	return c.stretgy.DoOperation(a, b)
}

func main() {
	c := &Calculator{}
	c.SetStrategy(AddOperation{})
	fmt.Println(c.calculate(5, 4))

	c.SetStrategy(SubtractOperation{})
	fmt.Println(c.calculate(5, 4))
}
