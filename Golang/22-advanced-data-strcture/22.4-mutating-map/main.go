package main

import "fmt"

/*

 */

func change(m map[string]int) {
	m["apple"] = 50
}

func main() {
	//var m map[string]int

	m := map[string]int{
		"apple":  1,
		"orange": 2,
	}

	fmt.Println(m)

	fmt.Println(m["apple"])
	m["apple"] = 20
	fmt.Println(m["apple"])

	if val, ok := m["apple"]; ok {
		fmt.Println("old value :", val)
		m["apple"] = 100

	}
	fmt.Println(m["apple"])

	delete(m, "apple")
	fmt.Println(m)

	change(m) //it is passed here as a reference underline pointer
	fmt.Println(m)
}
