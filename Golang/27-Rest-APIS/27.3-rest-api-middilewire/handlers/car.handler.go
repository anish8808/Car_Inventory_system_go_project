package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func carHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	entity := strings.TrimPrefix(path, "/cars")
	entity = strings.Trim(entity, "/")

	switch r.Method {
	case "POST":
		if entity == "" {
			createCar(w, r)
		} else {
			http.Error(w, "Incorrect post request", http.StatusBadRequest)
		}
	case "GET":
		if entity == "" {
			http.Error(w, "Incorrect get request", http.StatusBadRequest)
		} else {
			id, _ := strconv.Atoi(entity)
			getCar(w, r, id)
		}
	case "DELETE":
		if entity == "" {
			http.Error(w, "Incorrect delete request", http.StatusBadRequest)
		} else {
			id, _ := strconv.Atoi(entity)
			deleteCar(w, r, id)
		}
	}
}

var mu sync.Mutex

func createCar(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, " json format is woring ", http.StatusBadRequest)
		return
	}
	id := rand.Intn(10000)
	car.ID = id
	Cars[car.ID] = car

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

func getCar(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()
	car, ok := Cars[id]

	if !ok {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	fmt.Println("car found with the Id :", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(car)
}

func deleteCar(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()
	_, ok := Cars[id]

	if !ok {
		http.Error(w, "Delete request is not valid , id please check", http.StatusNotFound)
		return
	}
	delete(Cars, id)
	fmt.Println("the Id is found and deleted :", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
