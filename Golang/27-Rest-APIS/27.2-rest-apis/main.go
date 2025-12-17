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
	mu.Lock()
	defer mu.Unlock()
	var car Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, "Incorecct Json Body", http.StatusBadRequest)
		return
	}

	id := rand.Intn(1000)
	Cars[id] = car
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

func getCar(w http.ResponseWriter, id int) {
	mu.Lock()
	defer mu.Unlock()
	car, ok := Cars[id]
	if !ok {
		http.Error(w, "Incorrect rest call", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(car)
}

func deleteCar(w http.ResponseWriter, id int) {
	mu.Lock()
	defer mu.Unlock()

	_, ok := Cars[id]
	if !ok {
		http.Error(w, "Incorrect delete call , car is not present", http.StatusNotFound)
		return
	}

	delete(Cars, id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/cars", carHandeler)
	http.HandleFunc("/", carHandeler)
	fmt.Println("Server Started to listen the API Request")
	http.ListenAndServe(":3034", nil)
}
