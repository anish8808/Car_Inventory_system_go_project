/*
	- fixed size collection of element
	- creates a deep copy (value type)


	Decleration:
		- var arr[3] int  ---> [n]T
		- literal intialization
			arr:=[3]int{1 ,2 ,3}

		- multi dimessional array
			var arr[2][2]int

Slices
	- flexible dynamic views over arrays (views means pointe to an array)
	- it ca be grow and shrink

	internals of slices
		- pointer to an array
		- length (actual available element in the array)
		- capacity (based on the length its generate the capacity like vector only )
		- dynamic array

	Declearations:
		var s[] int //[]T
		s = arr[1:4]

		using make : we can create any particular object in golang
		s:= make([]int , 3 , 5)(lenght , capacity)

		Expansion :
			s:= []int {1 , 2 , 3}
			s= append(s , 4 , 5) //(name , variadic form ...)


*/

package main

import "fmt"

func incr(a [3]int) {
	a[0] = 100
}

func incrs(a []int) {
	a[0] = 100
}

func main() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Printf("A: %v", arr)

	B := [3]int{1, 2, 3}
	fmt.Printf("B: %v", B)

	var s []int
	s = append(s, 1, 2, 3) // 1 , 2 ,4 ... as variadic (also append doesnot return the pointer its generalised way to add element )
	fmt.Printf("s: %v", s)

	//this will work like call by value
	incr(B)
	fmt.Printf("A: %v", B)
	//this will work as call by Refrence
	incrs(s)
	fmt.Printf("s: %v", s)

	fmt.Printf("A: %v", len(B))
	fmt.Printf("A: %v", len(s))

}
