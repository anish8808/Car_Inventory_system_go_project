// fixed value and can not be change during execution

//used to define the value of compilation time

//declare is  used const name {type} = values(can not use short hand declearaton)
// const Pi float64 = 3.14

// constant Type
// 1- Typed (above example is types or u cant use typed)
// 2- Untyped
// consant we can define with block as well
// value can not be dynamic for constant
// Enumatared Constant with iota like
/*
	const (
	Apple = iota
	bananna
	cherry
	)
*/

// type conversation also support but changing value dont support
package main

import "fmt"

func main() {

	const a = 10
	const Pi float64 = 3.14

	fmt.Printf("a= %v , Type a = %T\nPi= %v  , Type of Pi = %T\n", a, a, Pi, Pi)
	const (
		L = 5
		W = 10
		A = L * W
	)
	fmt.Printf("Area = %v , Types = %T\n", A, A)
	const (
		Apple = iota
		bananna
		cherry
	)
	fmt.Println(Apple, bananna, cherry)
	z := float64(a)
	fmt.Printf("tyoe is %T", z)
}
