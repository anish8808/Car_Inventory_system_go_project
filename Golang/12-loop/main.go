/*
	#for loop :
		basic c styled for loop

		for i:=0; i<5; i++ {
			print(i)
		}

	#for as while loop  :
		i:=0
		for i<5{
			print(i)
			i++
		}

	// if no condition then it will be infine loop

	#for {
		print(wow)
	}

	#Range based

	numbers:= []{1 ,2 , 3, 4}
	for idx , val := range numbers {
		print(idx , val)

		// if you dont want to use anyone this idx or val where we can put the _ (this will ignore that value)
	}

	#OR

	for idx := range numbers {
		print(idx) // helpful for map

	}

	#if else condition

		#basic

			if x >5 {
				println("x>5")
			}


		# if else chain

			if x>5 {
				print(x)
			} else if {
				print(x)
			} else {
				print(x)
			}

		#if with short statements

			if y:=10; y>9 {
				print("wow")
			}

	#Switch Case
	//braek is not needed
	#Basic :

		much more powerful than other language

		day :=3
		switch day {
		case 1:
			print(1)
		case 2:
			print(2)
		case 3:
			print(3)
		default:
			print(other day)

		}

	#Mulitple values in switch case
			we cannot use the value in other other cases
	day :=3
		switch day {
		case 1 , 2 , 3:
			print(1)
		case 4 ,5:
			print(2)
		case 5 , 6:
			print(3)
		default:
			print(other day)

		}

	#type Switch

	var x interface {} =7
	switch x.(type){
		case int:
			//
		case string :
			//
		case float64
			//
		default:
			//
	}

*/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Basic loops

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//as while loop
	x := 5
	for x < 10 {
		fmt.Println(x)
		x++
	}

	x = 0
	for {
		if x > 10 {
			break
		}
		fmt.Println("infinete loop")
	}

	//Range for loop
	numbers := []int{1, 2, 3, 4, 5}

	for index, val := range numbers {
		fmt.Println(index, val)
	}

	//IF ELSE CONDITION

	//basic
	if x >= 5 {
		fmt.Println("x is greater then 5")
	}

	if x >= 5 {
		fmt.Println("x is greater then 5")
	} else {
		fmt.Println("x is lesser then 5")
	}

	//if with short statement

	if _, e := strconv.Atoi("124"); e != nil {
		fmt.Println(e)
	}

	//Switch case

	day := 3

	switch day {
	case 1:
		fmt.Println("its day 1")
	case 2:
		fmt.Println("its day 2")
	case 3:
		fmt.Println("its day 3")
	default:
		fmt.Println("default")
	}

	//multiple cases
	switch day {
	case 1, 2:
		fmt.Println("its day 1")
	case 3, 4:
		fmt.Println("its day 2")
	case 5:
		fmt.Println("its day 3")
	default:
		fmt.Println("default")
	}

	//typed switch

	var i interface{} = 7

	switch i.(type) {
	case int:
		fmt.Println("i is int")
	case float64:
		fmt.Println("its float")
	default:
		fmt.Println("i dont know")
	}

}
