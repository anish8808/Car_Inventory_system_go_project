package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"sync"
)

type Car struct {
	Id    int64
	Name  string
	Model string
	Brand string
	Yeatr int64
	Price float64
}

// {
// 	"name": "X7",
// 	"model": "v6 inline",
// 	"brand": "BMW",
// 	"year": 2024 ,
// 	"price": 20000000
// }

var Cars = make(map[int64]Car)
var mu sync.Mutex

func carHandeler(r *http.Request, w http.ResponseWriter) {

	path := r.URL.Path

	entity := strings.TrimPrefix(path, "/cars")
	entity = strings.Trim(entity, "/")

	switch r.Method {
	case "POST":
		if entity == "" {
			//create handelr
		} else {
			http.Error(w, "Incorect post Request", http.StatusBadRequest)
		}
	case "GET":
		if entity == "" {
			//create handeler
		} else {
			http.Error(w, " We dont Support this API ", http.StatusBadRequest)
		}
	case "PUT":
		if entity == "" {
			http.Error(w, " We dont Support this API ", http.StatusBadRequest)
		} else {
			//put handler
		}
	}
}

func createCar(r *http.Request, w http.ResponseWriter) {
	mu.Lock()

	defer mu.Unlock()

	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "incorrect json input", http.StatusBadRequest)
	}

	id := rand.Intn(10000)
	car.Id = int64(id)
	//store in the map
	Cars[car.Id] = car

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader((http.StatusAccepted))
	json.NewEncoder(w).Encode(car)
}

func getCar(r *http.Request, w http.ResponseWriter, id int64) {
	mu.Lock()

	defer mu.Unlock()

	car, ok := Cars[id]
	if !ok {
		http.Error(w, "Not found car", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader((http.StatusOK))
	json.NewEncoder(w).Encode(car)
}
func main() {

}
