/*
	DEFER :
		- defer is a keyword that delays the executions of the function or method ,or anyomnus function
		- evaluates instantly , execute at the end just before your function out of stack or exiting
			weather your function closed normally , or with some panic or with some error,
			defer will always get executed before the return statement of the program

		- multiple statements are allowed
			- it will work as LIFO

		SYNTAX:\

		func surroundings (){
		defer GFG() -- this will run at the end of this function
		//code blocks
		}

		//example of multiple defer
		func surroundings (){
		defer GFG() -- this will run at the end of this function
		defer GFG2() --> this will execute at last but first to gfg()
		//code blocks
		}

		IMP ** -- (it has overhaed memoery as stack has to use to mainain the call )
*/

package main

import (
	"fmt"
	"os"
)

func demo(i int) {
	fmt.Println("Hello GFG", i)
}

func counters() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}
func main() {
	defer demo(1) //--> third and last run

	demo(2) //-->first run

	defer demo(3) //-->second run

	counters()

	//3rd example Real world
	file, err := os.Open("temp.txt")
	if err != nil {
		fmt.Errorf("error is there %v", err)
	}

	defer file.Close() // this will close the opened file at the end of the your program
}
