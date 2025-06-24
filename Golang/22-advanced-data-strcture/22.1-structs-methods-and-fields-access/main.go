package main

import "fmt"

/*
	Structs : method and fields access
		allows you to define behaviour and manipulate the data defined in the strctures
		fields access using (.) operator

	method on strctures:
		-->function that are associated with strctures type called method
		-->method can be defined with recivers -- essentially its a strctures on which you are defining the function
			-->two types of recivers in golang
					1-value: define strcture and passed by value(actually passing the value of strcture and passing a copy)
						by using that we can not actual value in the original strctures , only change will be in functional scope
					2- Pointer : define strctures and passed by pointer
*/

type car struct {
	name  string
	Type  string
	Brand string
	years int
}

//defining the methods on car

func (c car) operational() {
	if c.years > 15 {
		fmt.Printf("car with: %s and brand %s is not operational.\n", c.name, c.Brand)
	} else {
		fmt.Printf("car with: %s and brand %s has %v years of operational left.\n", c.name, c.Brand, 15-c.years)
	}
}

func main() {
	c := car{"Modes", "EV", "Tesla", 3}
	fmt.Println(c)
	c.operational()
}
