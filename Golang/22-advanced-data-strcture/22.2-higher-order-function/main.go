/*
functions are first class citizens
	Higher order funs and function closure
		1- Higher-order-function:
			--> function takes one or more functions as aurgement
			--> return a function
		2- inserting a fun as paramtere
		func take input as a function:
			func operation(x , y int  , op func(int , int) int) int{
				return op(x , y)
			}

		fun add(x , y int) int{
			return x+y
		}

		\\--> now this func we can use in operation singnature

		int ans := operation (2 , 3 , add)

		3 - return function

		func muntiplier (factor int ) func(int)int {
			return func(x int) int{
				return x*factor
			}
		}
*/

package main

import (
	"fmt"
)

func operation(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func add(x, y int) int {
	return x + y
}

func muntiply(x, y int) int {
	return x * y
}

func muntiplier(factor int) func(int) int {
	return func(base int) int {
		return base * factor
	}
}
func main() {

	//1. function as input
	fmt.Println(operation(2, 3, add))
	fmt.Println(operation(2, 3, muntiply))

	mult := muntiplier(3)
	fmt.Println(mult(4))
}
