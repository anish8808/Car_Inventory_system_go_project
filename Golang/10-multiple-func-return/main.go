/*
Multiple Result:
	go has a distinguishing Feature ability to return the multpile values
	useful while returning the error along with the result


	func name (p1 t1 , p2 t2) (return t1  , return t2){
		//function body
		return value1 , value2
	}

*/

package main

import (
	"fmt"
	"strconv"
)

func swap(x, y int) (int, int) {
	return y, x
}

func Convert(str string) (int, error) {
	return strconv.Atoi(str)
}
func main() {
	fmt.Println(swap(2, 4))

	i, e := Convert("123") //valid integer  here
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(i)

	i, e = Convert("xyasa") // invalid integer
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(i)

}
