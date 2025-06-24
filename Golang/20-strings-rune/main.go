/*
	Stirgs:
		Read only Sequence of bytes
		Inmutable : once created , cant be changes
		UTF-8 encoded (each char 1-4 bytes) -- > means you have a special unicode character which can take
			upto 4 bytes to represent one character like emoji etc .(its not single char like a,  b ,c ) these
			are sets of character which make UTF-8 incoded character  , that is the reason when you have set of
			bytes in your strings if we break in between you may loose some unicode charcater i.e run come into the picture


	Declearations:
		var str1 string ="Hello Gfg"
		var str2 string = `hello
							gfg`

	propersties:
		length = len(str1)
		greetings := "hello" + "Gfg"
		access the char based on the index
		for i , ch := range str1{
		fmt.println("%c" , ch)
		}

	conversations can :
		b := []byte("hello")
		b[0]= 'H'


Rune :
	--> alias for int32
	--> strings are sequence of bytes but runes are sequence of unicode character
	--> runes are essantially represenataion of single UTF 8 or unicode character
	-->lets have the emoji , set of bytes will be 0 to 3 bytes for emoji , so this character
		can be easily understand y the golang using runes
	-->

	Declarations :
		var r rune = 'emoji'
		print(r)

	conversations:
		slice := strings = seq of bytes
				[] runes = seq of unicode char

		s: = "hello, emoji"
		r:= []rune(s)
		each element in r is unicode char

		--> go uses UTF encoding
		--> each unicode char take 1 to 4 bytes
		--> ASCII = 1 byte , chinese = 3/4 bytes



why runes are important :
	--> Beacuse lets say we have some https based server which needs to respond in diff diff languages
	in case of thr localiazations , in that case we can not just return an array of string , we have to
	convert into the rune , and rune will take care of the all unicode character and localization will be
	easier or similar response will go into the english and similiar conversion will go in chiness or etc
*/

package main

import (
	"fmt"
)

func main() {
	s := "Hello GFG! 😃"
	r := []rune(s)
	fmt.Println(len(s))
	fmt.Println(len(r))

	for i, ch := range s {
		fmt.Printf("String char %c , and index %d , with the ASCII %v\n", ch, i, ch)
	}
	fmt.Printf("rune print \n")
	for i, ch := range r {
		fmt.Printf("Rune char %c , rune index %d , with the ASCII %U\n", ch, i, ch)
	}

	s = s + "Cool"
	fmt.Println(s)
}
