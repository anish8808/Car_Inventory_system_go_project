package main

import "fmt"

/*
	// provides ability to encapsulate the code to reuse


	func Funcationname (p1 t1 , p1 t2 , ...)(r1 t1 , r2 t2){

	function body
	return statement
	}



Variadic function
	-> accepts varibale number of aurgements (if you dont know how many parmaters to send)

	func sum (numbers ...int --> this will be slice) int {
		//total sum(range over the slice and add to sum all the numbers in the slice)
		//return total
	}
	--> we can call sum(1 ,2 3 ,) or sum (1 , 2)


Anonyummus functions:  when you want to assgin a function to a varibale we can use anonymous fucntion
this is like inline function support for each variable

greet := func (name string) {  // we are not having anyname
	fmt.Println("Hello" , name)
}

greet("GFG")

*** Function closure : -> special type of anonymous functions , it references variable outside of its scope
-->benifits , will be able to access the varibale outside of the scope
*/

func Hello(a int, b int) int {

	fmt.Println("a+b", a+b)
	return a + b
}

func Hello2(a, b int) int {

	fmt.Println("a+b", a+b)
	return a + b
}

//variadic function

func Add(num ...int) int {
	total := 0
	for i := range num {
		total += num[i]
	}

	return total
}

func Counter(a int) func() int {
	count := 10

	return func() int {
		count = count + a // this function us able to access the count variable even it is the out of the scope
		return count
	}
}

func main() {
	ans := Hello(2, 3)
	fmt.Println("ans", ans)
	ans = Hello2(2, 3)
	fmt.Println("ans", ans)

	// varadic function
	ans = Add(1, 2, 4)
	fmt.Println("ans", ans)

	ans = Add(1, 2, 4, 10, 12)
	fmt.Println("ans", ans)

	//anonyums functions
	greet := func(name string) {
		fmt.Println("anonymus function", name)
	}

	greet("ANISH using GFG")

	count := Counter(3)
	fmt.Println(count())

}
