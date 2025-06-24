package main

import (
	"errors"
	"fmt"
)

/*
Errors in go
	Go takes unique approaches towards error handling
	Returns errors as values insted of exception(like other programming language)


	1- creating an error:(builin package in go errors)
		func check(num int) (string , error){  (//error is herr variable that is used by error package)
		 if num<0 {
		 	return " " , errors.New("negative n")
		 }
		 return "Positive num" , nil (--> in error 0 means nil)
		}

	2- wrapping of errors: ( in any format )
		err:=check(-1)
		if err != nil {
			return fmt.Errorf("failuer: %w" , err)
		}
*/

func check(num int) (string, error) {
	if num < 0 {
		return " ", errors.New("Number is negative")
	}

	return "Number is not negative", nil
}

func read(filename string) error {
	if filename == "" {
		return ErrinvalidFile
	}
	return nil
}

func file() (string, error) {
	e := read("")
	if e != nil {
		return "", fmt.Errorf("Error reading file :%w", e)
	}
	return "file exists", nil
}

var ErrinvalidFile = errors.New("File is not valid")

func main() {
	//basic error handling
	_, e := check(-1)
	if e != nil {
		fmt.Println(e)
	}

	//wrpaping of error

	_, err := file()
	if err != nil {
		fmt.Println(err)
	}

	//unwrpaping of error
	we := errors.Unwrap(err)
	fmt.Println(we)

	if errors.Is(we, ErrinvalidFile) { // this is check for two errors (but is always unwrap then comaper only)
		fmt.Println(we)
	}
}
