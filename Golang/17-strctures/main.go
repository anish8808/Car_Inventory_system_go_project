/*
	- combine the diff diff types of data types to make meaningful use
	- structs are user defined types that group fields togather
	- collection of diff diff data types
	- same as anonymus strcut there is no name for struct with short hand declerations
	- Strctures Embending - another structures in same strctures
	- go allows to check weather 2 structs same or not

*/

package main

import "fmt"

func main() {

	type Power struct {
		HP   int
		Type string
	}
	type Car struct {
		Name  string
		Type  string
		Brand string
		years int
	}

	type Car2 struct {
		Power
		Brand string
		years int
	}
	var c Car
	c = Car{Name: "Model s", Type: "EV", Brand: "Tesla", years: 5}
	fmt.Println(c.Name)

	var c2 Car2
	c2 = Car2{Power: Power{HP: 500, Type: "supera"}, Brand: "new", years: 10}
	fmt.Println(c2.Brand)
	//positional struct filling but in this case we have to consider all the variables of the strctures

	c = Car{"Modes", "ev", "tesla", 3}
	fmt.Println(c.Name)

	//named field

	c = Car{Type: "Ev"} // in this case all varibles will be zero

	//anonymus strctures
	p := struct {
		Name string
		age  int
	}{
		Name: "Anish",
		age:  25,
	}

	fmt.Println(p.Name, p.age)

}
