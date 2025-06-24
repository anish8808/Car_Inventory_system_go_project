package main

import "fmt"

/*
Interfaces :

  - collection of Methods signatures

  - Any type that implemented all methods satisfy the interface

  - no need of explicit declaration like implements or extends in java

    type Car Interface {
    startEngine() string
    Drive() string
    }

    Implemeting the interface :
    type Sedan Struct {

    Brand string
    Length int
    }

    func (s Sedan) StartEngine() string {  //matched the signature of interface
    }

    func (s Sedan) Drive() string {  ///matched the signature of interface
    }

	** How to assign the dyanmic types to an interface
		we can only assign such types which implements the signatature method that are defined Car interface
		now with single interface we can assign multiple types dynamically

	** PolyMorphism :
		here by using single car interface we are able to print differnet differnet types of struct
		this is a way to achive the polymorphism
*/

type Car interface {
	StartEngine() string
	Drive() string
}

func DefineCar(c Car) {
	fmt.Println(c.Drive())
	fmt.Println(c.StartEngine())
}

type Sedan struct {
	Brand  string
	Length int
}

type SUV struct {
	Brand  string
	Height int
}

func (s Sedan) StartEngine() string {
	return fmt.Sprintf("%s Brand car has stared  engine silently\n", s.Brand)
}

func (s Sedan) Drive() string {
	return fmt.Sprintf("%s Brand sedan has length %dfeet can be driveable smoothly on highways", s.Brand, s.Length)
}

func (s SUV) StartEngine() string {
	return fmt.Sprintf("%s Brand car has stared  engine silently\n", s.Brand)
}

func (s SUV) Drive() string {
	return fmt.Sprintf("%s Brand sedan has length %dfeet can be driveable smoothly on highways", s.Brand, s.Height)
}

func main() {
	var car Car

	car = Sedan{Brand: "Honda", Length: 15}
	DefineCar(car)
	car = SUV{Brand: "Jeep", Height: 20}
	DefineCar(car)
}
