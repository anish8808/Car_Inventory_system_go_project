/*
	similar to c
	var p* T(this can be anything)
	pointer assignment :
		var x int =10
		p:=&x , // p = address of x
	derefrencing
	y = *p

	use case:
		- manage memories efficiently
		- pass by refrence modifing the values in functions
		- go doest not allow pointer airthmatic like c and c++
		- initaial var p* int will be nil
*/

package main

import (
	"fmt"
)

func incr(x *int) int {
	return *x + 1
}
func main() {
	var p *int
	x := 42
	p = &x
	fmt.Println(p, x)

	*p = incr(p)
	fmt.Println(p, x)

}
