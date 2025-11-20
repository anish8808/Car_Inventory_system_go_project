package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

/*
	CAR INVENTORY SYSTEM
	1 - car information like car dekho app
		- ID of the car
		- name of the CAR ,
		- years of manufacturing
		- price


	2 - list of CAR
		- Add a car to an inventory
		- Get the information of a car from the inventorty
		- Delete the car from the inventory


*/

type Car struct {
	ID    int64
	Name  string
	Model string
	Brand string
	Year  int64
	Price float64
}

var Cars = make(map[int64]Car)

var mu sync.Mutex

func carHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	entity := strings.TrimPrefix(path, "/cars")
	entity = strings.Trim(entity, "/")

	switch r.Method {
	case "POST":
		if entity == "" {
			createCar(w, r)
		} else {
			http.Error(w, "Incorrect post Request ", http.StatusBadRequest)
		}
	case "GET":
		if entity == "" {
			http.Error(w, "we dont support this API ", http.StatusBadRequest)
		} else {
			id, _ := strconv.Atoi(entity)

			getCar(w, int64(id))
		}
	case "DELETE":
		if entity == "" {
			http.Error(w, "we dont support this API", http.StatusBadRequest)
		} else {
			id, _ := strconv.Atoi(entity)

			deleteCar(w, int64(id))
		}
	}
}

/*
JSON Body


{
	"name": "X7"
	"Model": "v6 inline"
	"Brand": "BMW"
	"Year": 2024
	"price": 200000000
}
*/

func createCar(w http.ResponseWriter, r *http.Request) {

	mu.Lock()

	defer mu.Unlock()

	var car Car

	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Incorrect json input", http.StatusBadRequest)
		return
	}
	id := rand.Intn(10000)

	car.ID = int64(id)

	Cars[car.ID] = car //storing in a database

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusCreated))
	json.NewEncoder(w).Encode(car)
}

func getCar(w http.ResponseWriter, id int64) {
	mu.Lock()

	defer mu.Unlock()

	car, ok := Cars[id]

	if !ok {
		http.Error(w, "car not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(car)

}

func deleteCar(w http.ResponseWriter, id int64) {
	mu.Lock()
	defer mu.Unlock()

	_, ok := Cars[id]

	if !ok {
		http.Error(w, "Car info is not present ", http.StatusNotFound)
		return
	}

	//delete map valuet
	delete(Cars, id)

	w.WriteHeader(http.StatusNoContent)

}

func main() {
	http.HandleFunc("/cars", carHandler)
	http.HandleFunc("/cars/", carHandler)

	fmt.Println("HTTP server is listing...")
	http.ListenAndServe(":3015", nil)
}
