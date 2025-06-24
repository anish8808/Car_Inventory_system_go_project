package main

import (
	"fmt"
	"strconv"
)

//it is procress of converting from onr type to another type
//explicit always
// target := T(V) {type , variable}

func main() {

	//integer to float
	i := 5
	f := float64(i)
	fmt.Println("int to float64:", f)

	//float to integet
	k := 0.1
	in := int(k)
	fmt.Println("int to int:", in)
	//string to slice
	s := "anish"
	b := []byte(s)
	fmt.Println("int to slice :", b)
	//string to int

	new := []byte{1, 2, 3}
	str := string(new)
	fmt.Println("int to string:", str)

	// interget to string
	s = strconv.Itoa(i)
	fmt.Println("int to string:", s)

	//string to int
	i, _ = strconv.Atoi(s)
	fmt.Println("convrt ", i)
}
