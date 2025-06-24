package main

import "fmt"

/*
	Panic and recover is similar to try and catch like exception heandling
	Panic :
		   Used when something goes unexpectedly wrong(addressing a memory that is the nil pointer or division by zero)
		   Panic is for unrecoverable situations
		   stops the normal flow of execution & begins panicking with some log


		func divide (a ,b int) int {
		if b==0{
			panic("division by zero")
		}
		return a/b
		}


	Recover :
		recover is used to regain control of panicking goroutines
		only works when called using defer function
		it catches panics & prevent crashes in critical parts

	func main() {
		defer func (){
			if r:= recover(); r!= nil {
				fmt.println(r) --> this will recover the panic
			}
		}()

		divide(4 , 0)
	}

*/

func divide(a, b int) int {
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}
func main() {

	fmt.Println(divide(4, 2))
	fmt.Println(divide(4, 0)) //--> program will terminate here and furture noet exxcute to next line
	fmt.Println(divide(4, 1)) // this is blocked ( for recover )
}
