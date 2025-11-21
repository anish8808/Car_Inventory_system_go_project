package main

import (
	"car/config"
	"car/handlers"
	"fmt"
	"net/http"
)

func main() {

	config.ConnectDB()
	http.HandleFunc("/cars", handlers.CarHandler)
	http.HandleFunc("/cars/", handlers.CarHandler)

	fmt.Println("Serever started running and listen")
	http.ListenAndServe(":3015", nil)
}
