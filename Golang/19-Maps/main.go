/*
	Powerful data strctures that implement hash table
	fast lookups
	O(1) time complexity

	Declearations :
		var m map[string]int  // var name map[T1]T1/T2/T3
		- initally value in the map will be nil
		insertion:
			map["key"] = 42
		initalizations:
			m:=make(map[string]int)
		accessing :
			val := map["GFG"]
			val ,ok := map["foo"] , if key is not present then the value of the val will be Zero of the data type
		delete:
			delete(map , "GFG")
*/

package main

import "fmt"

func chnage(m map[int]string) {
	m[100] = "hello Gfg"
}

func main() {
	var m map[int]string

	if m == nil {
		fmt.Println("map is nil")
	}

	m1 := make(map[int]string)

	m1[100] = "GFG"

	val := m1[100]
	fmt.Println(val)

	val, exists := m[200]

	if !exists {
		fmt.Println("map key doest exist")
	} else {
		fmt.Println("map key  exist")
	}

	chnage(m1) //map is also a refrence type data strctures in golang
	val = m1[100]
	fmt.Println(val)

}
