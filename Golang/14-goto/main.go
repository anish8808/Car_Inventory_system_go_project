/*
	GOTO:
		--> flow control statement that jumps to  labelled statement withing the same function
		--> not much encourged to use it


	BASIC USAGE:

	func demo (){
		print(n)
		goto end
		print(n)
	end:
		print(n)
	}

	-->avoid GOto wherever possible
	--> it creates confusing code (spaghetti code)
*/

package main

import "fmt"

func demo() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				goto end
			}
			fmt.Println(i, j)
		}
	}
end:
	fmt.Println("Exit of the loops")
}

func main() {
	demo()
}
