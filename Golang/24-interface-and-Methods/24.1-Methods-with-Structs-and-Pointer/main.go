package main

import "fmt"

/*
	Methods with strctures & Poniters
		1 - Methods can be associated with both
		2 - allows cleaner & More encapsulated data

		Strcut/Pointer Reciver Types :

		Strcut Reciver Types                                       Pointer Reciver Types
		Copy of struct is passed                                   Refrence of struct is paased
		Any changes in the method wont affect original struct      Method can modify original Struct


		example for both

		Type Rectangle struct {
			width int
			Height int
		}
		1 - strcut reciver Types :

			fun( r Rectangle ) Area () int {
				return r.width * r,height  // here copy of the struct
			}

		2 - func (r * Rectangle) scale (f int) {
			r.width * = f
			r.height * = f
		}
*/

type Rectangle struct {
	width  int
	hieght int
}

func (r Rectangle) Area() int {
	return r.width * r.hieght
}

func (r *Rectangle) scale(f int) {
	r.width *= f
	r.hieght *= f

}

func main() {
	rec := Rectangle{width: 10, hieght: 10}
	fmt.Println(rec.Area())
	rec.scale(10) // this will scale the original Rectangle
	fmt.Println(rec.Area())
}
