package main

import (
	"fmt"
)

/*
Type assersation & Type switches :

	-> interface {} can hold any type in go
	-> flexibility when working with dynamic types
	-> helps when need to extract concrete type
	-> two component are stored
		1- actual concrete
		2- type of the value

Example :
Type assersation:

	value , ok := interfaceValue.(ConcreteType)

	var v interface{}
	v = 45
	num := v.(int)
	v = "GFG"
	str := v.(str)

Type switches :

	switch v:= interfaceValue.(type){
		case concreteType1:
			Logic1
		case concreteType2:
			Logic2

	}
*/
type Car interface {
	StartEngine() string
	Drive() string
}

func DefineCar(c Car) {
	//--> Implementing the type switches are here
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

func TestDrive(c Car) {
	switch tp := c.(type) {
	case Sedan:
		tp.Drive()
	case SUV:
		tp.Drive()
	default:
		fmt.Println("we dont have other types of car")
	}
}

func main() {
	var output interface{}
	output = 45
	i, ok := output.(int)

	if !ok {
		fmt.Println("Unable to convert output ")
	}
	fmt.Printf("Integer %d", i)

	output = "GFG"
	str, ok := output.(string)

	if !ok {
		fmt.Println("Unable to convert output")
	}
	fmt.Printf("String %v", str)

	switch tp := output.(type) {
	case int:
		fmt.Println(tp, " as integer")
	case string:
		fmt.Println(tp, " as string")
	default:
		fmt.Println("unsuported output ")
	}

	var c Car
	c = Sedan{Brand: "Honda", Length: 20}
	TestDrive(c)
	c = SUV{Brand: "Jeep", Height: 20}
	TestDrive(c)
}
