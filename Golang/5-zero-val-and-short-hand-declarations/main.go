package main

import "fmt"

// in Go, when a variable is declared without any initial value,
// it always has a predictable zero value

// 1- int, int8, int64 == 0
// 2- float32, float64 == 0.0
// 3- bool == false
// string == ""
// arrays, slices = [0 0 0], [] // value based on the data types
// maps = map[] (empty map)
// pointers, interface = nil

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var arr [3]int
	var slc []int
	var m map[string]int
	var p *int
	var iface interface{}

	fmt.Println("int:", i)
	fmt.Println("float64:", f)
	fmt.Println("bool:", b)
	fmt.Println("string:", s)
	fmt.Println("array:", arr)
	fmt.Println("slice:", slc)
	fmt.Println("map:", m)
	fmt.Println("pointer:", p)
	fmt.Println("interface:", iface)
}
