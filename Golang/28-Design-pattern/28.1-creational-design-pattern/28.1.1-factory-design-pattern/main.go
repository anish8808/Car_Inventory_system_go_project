package main

import "fmt"

/*
	factory design pattern is a crational desing pattern , and one of the most commonly used dessing patter
	this pattern provides a way to hide the creation of the logic of the instance
	the client only interacts with a factory and tell the kind of instance that needs to be created
	then the factory interfaces with the crroesponding concrete objects and return s the correct instance back

	# class creation should be abstracted from the client , client is actually interacting with factory
	this is the benift so client and logic shound not exposs to each other
*/

type Car interface {
	getCar() string
}

type SedanCar struct {
	Name string // this is the app code
}

func (s *SedanCar) getCar() string {
	return "Honda city"
}

func getNewSedan() *SedanCar {
	return &SedanCar{}
}

type hatchBack struct {
	Name string
}

func (h *hatchBack) getCar() string {
	return "Nexon"
}

func getNewHatchBack() *hatchBack {
	return &hatchBack{}
}

func getCarFactory(cartype int) { // factory method
	var car Car
	if cartype == 1 {
		car = getNewHatchBack()
	} else {
		car = getNewSedan()
	}

	carInfo := car.getCar()
	fmt.Println(carInfo)
}

func main() { //client method
	getCarFactory(1)
	getCarFactory(2)
}
