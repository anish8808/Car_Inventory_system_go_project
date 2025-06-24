package main

import "fmt"

/*
Custom Error strcture

	type error interface {
		Error() string
	}
*/

type CustomError struct {
	code    int
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Error code %d with message %s", e.code, e.Message)
}

func check(code int) (string, error) {
	if code == 400 {
		return "", &CustomError{code: 400, Message: "Invalid request"}
	}
	return "success", nil
}

func main() {
	if _, err := check(400); err != nil {
		fmt.Println(err)
	}
}
