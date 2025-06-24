package main

import (
	"fmt"
)

/*
	Closures :
		oftern used in context of returning of a function (higher ordere)
		closure captures the environment in which it was defined
		even after outer function returns clouser retains the access of the variables out scope
		it works with Go's lexical scoping rules (surrounding varibale of a func will be avaible once you return the function)


	func add() func(int) int {
		x:=5
		return fun(i int ) int {
			return x + i
		}
	}

	a:= add() --> it will be access z variable a here will be treated as func
	a(3) -- > here function closuring best example
*/

func adder() func(int) int {
	x := 5
	return func(i int) int {
		return x + i
	}
}

func fun() (func(), func()) {
	x := 0

	inc := func() {
		x++
		fmt.Println(x)
	}

	dec := func() {
		x--
		fmt.Println(x)
	}

	return inc, dec
}
func main() {
	a := adder()
	fmt.Println(a(5))

	in, de := fun()
	in()
	de()
	in()
	de()

}
