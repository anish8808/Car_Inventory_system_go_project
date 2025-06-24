/*
	Named result values
		Go Allows naming of the return values
		Simplifies function body
		auto matching gets intialized to zero values

		func Swap (x , y int) (a int , b int){
		a =y
		b =x
		return // implicitly go understand return a , and b
		}

	//Benifits is enhance readibility
*/

package main

import "fmt"

func swap(x, y int) (a, b int) {
	a = y
	b = x
	return
}

func area(a, b float64) (area float64, p float64) {
	area = a * b
	p = 2 * (a + b)
	return
}

func main() {
	x, y := swap(5, 6)
	fmt.Println(x, y)

	ar, p := area(5, 6)
	fmt.Println(ar, p)
}
