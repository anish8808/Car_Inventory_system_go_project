package main

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Car struct {
	ID    int
	Name  string
	Brand string
	Year  int
	Model string
	Price int
}

var Cars = make(map[int]Car)

var mu sync.Mutex

func carHandeler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	entity := strings.Trim(path, "/cars")
	entity = strings.Trim(entity, "/")

	switch r.Method {
	case "POST":
		{
			if entity == "" {
				createCar(w, r)
			} else {
				http.Error(w, "Incorrect Rest Call ", http.StatusBadRequest)
			}
		}
	case "GET":
		{
			if entity == "" {
				http.Error(w, "Incorrecr Get rest request ", http.StatusBadRequest)
			} else {
				id, _ := strconv.Atoi(entity)
				getCar(w, id)
			}
		}
	case "DELETE":
		{
			if entity == "" {
				http.Error(w, "Incorrect delete request", http.StatusNoContent)
			} else {
				id, _ := strconv.Atoi(entity)
				deleteCar(w, id)
			}
		}
	}
}

func createCar(w http.ResponseWriter, r *http.Request) {

}

func getCar(w http.ResponseWriter, id int) {

}

func deleteCar(w http.ResponseWriter, id int) {

}

func main() {

}
