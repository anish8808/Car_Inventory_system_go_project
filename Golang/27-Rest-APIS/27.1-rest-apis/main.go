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
Rest APIS :
REST (Representational State Transfer) is an architectural style rather than a protocol,
widely adopted for designing networked applications. Introduced in 2000,
REST quickly became the standard for developing web services due to its simplicity and its ability to extend existing systems built around HTTP.

Key Characteristics of REST APIs
REST APIs revolve around a client-server architecture and are characterized by the following features:

Statelessness: Each request from a client to the server must contain all the information needed to understand and complete the request. The server does not retain any session information.
Client-Server Decoupling: The client and server are independent, allowing each to evolve separately without affecting the other.
Layered System: REST allows the use of a layered system architecture where intermediaries (such as proxy servers and gateways) can improve scalability and enforce security.
Uniform Interface: This simplifies the architecture, making interactions between components more predictable and understandable.
Cacheability: Responses must define whether they are cacheable or not, helping to prevent clients from reusing stale or inappropriate data.
The Role of HTTP in REST APIs
REST APIs use HTTP request methods to perform operations on resources. The primary HTTP methods include:

GET: Retrieve information about a resource.
POST: Create a new resource.
PUT: Update an existing resource.
DELETE: Remove a resource.
Responses from REST APIs are typically in JSON or XML format, with JSON being the more popular choice due to its easy integration with web applications and its simplicity for parsing.

Best Practices in REST API Design
When designing RESTful services, it's crucial to follow best practices to ensure the system is scalable, maintainable, and efficient:

Resource Naming: Use consistent, noun-based naming conventions. Collections should be named using plural nouns (e.g., /users).
Use HTTP Status Codes: Properly utilize HTTP status codes to communicate the outcome of API requests effectively. For example, 200 OK, 201 Created, 404 Not Found, and 500 Internal Server Error.
Versioning: Manage changes to the API properly through versioning (e.g., /v1/users), allowing clients to continue using an older version of the API and upgrade gradually.
Security: Implement strong authentication and authorization strategies to protect the API from unauthorized access (e.g., OAuth 2.0, JWT tokens).
Support for Hypermedia (HATEOAS): Include hypermedia links in responses to guide clients on possible next actions, which helps in decoupling client and server interactions.
Implementing REST API in Golang
Golang, with its robust standard library and powerful features for concurrent operations, is an excellent choice for implementing RESTful services. When developing a REST API in Golang, it is essential to:

Structure your application around the net/http package, which provides HTTP client and server implementations.
Utilize third-party libraries like Gorilla Mux for routing and middleware support to handle more complex scenarios than the standard library offers.
Ensure that your API handlers are concurrent-safe and capable of handling multiple requests efficiently.










HTTP methods and status codes.

comments
When designing RESTful APIs, understanding HTTP methods and status codes is fundamental. These elements create a standardized way for servers and clients to communicate effectively, ensuring smooth interactions and proper error handling. In this article, we’ll break down HTTP methods, their use cases, and status codes in a simple and precise manner.

HTTP Methods
HTTP methods define the type of operation to be performed on a resource. They are stateless and mostly idempotent (except for some exceptions like POST). Let’s explore the key methods:

1. GET
Purpose: Retrieve data from the server.
Idempotency: Yes, calling it multiple times does not change the resource state.
Use Case: Fetching resources, such as user information or map data.
Example:
GET /users – Fetch all users.
GET /users/{id} – Fetch a specific user by ID.
2. POST
Purpose: Create a new resource on the server.
Idempotency: No, repeated calls can result in duplicate data or conflicts.
Use Case: Submitting forms, creating user accounts, or adding entries to a database.
Example:
POST /users – Create a new user with details provided in the request body.
3. PUT
Purpose: Update or create a resource (if it doesn’t exist, based on API design).
Idempotency: Yes, calling it multiple times produces the same result.
Use Case: Updating an entire resource.
Example:
PUT /users/{id} – Update a user’s complete details by providing all fields.
4. PATCH
Purpose: Partially update an existing resource.
Idempotency: Yes, repeated calls update the same resource in the same way.
Use Case: Modifying specific fields of a resource.
Example:
PATCH /users/{id} – Update only the email or phone number of a user.
5. DELETE
Purpose: Remove a resource from the server.
Idempotency: Generally yes, calling it multiple times results in the same response (resource no longer exists).
Use Case: Deleting accounts, entries, or records.
Example:
DELETE /users/{id} – Delete a specific user.
6. HEAD
Purpose: Retrieve headers without the response body.
Idempotency: Yes.
Use Case: Check resource metadata (e.g., content length or type) without downloading the entire resource.
Example:
HEAD /users – Fetch metadata about the user collection.
7. OPTIONS
Purpose: Describe the communication options available for a resource.
Idempotency: Yes.
Use Case: Check supported HTTP methods for a resource.
Example:
OPTIONS /users – List methods like GET, POST, etc., allowed for the /users endpoint.
HTTP Status Codes
HTTP status codes indicate the result of a client’s request to the server. They are categorized into five classes:

1. Informational (1xx)
Purpose: Indicate that the request is received and being processed.
Common Codes:
100 Continue – The initial request is received, and the client can continue.
101 Switching Protocols – The server agrees to switch to a different protocol.
2. Success (2xx)
Purpose: Indicate successful processing of the request.
Common Codes:
200 OK – The request was successful, and the response contains the data.
201 Created – A new resource has been created (commonly used with POST).
204 No Content – The request was successful, but no content is returned (e.g., after DELETE).
3. Redirection (3xx)
Purpose: Indicate that further actions are needed to complete the request.
Common Codes:
301 Moved Permanently – The resource is permanently moved to a new URL.
302 Found – The resource is temporarily located at a different URL.
304 Not Modified – The resource has not been modified since the last request (used with caching).
4. Client Error (4xx)
Purpose: Indicate issues with the client’s request.
Common Codes:
400 Bad Request – The server could not understand the request due to invalid syntax.
401 Unauthorized – Authentication is required to access the resource.
403 Forbidden – The client is authenticated but lacks permission to access the resource.
404 Not Found – The requested resource does not exist.
409 Conflict – The request conflicts with the current state of the resource (e.g., duplicate entries).
5. Server Error (5xx)
Purpose: Indicate server-side issues.
Common Codes:
500 Internal Server Error – An unexpected condition occurred on the server.
502 Bad Gateway – The server acting as a gateway received an invalid response.
503 Service Unavailable – The server is overloaded or undergoing maintenance.
504 Gateway Timeout – The server acting as a gateway did not receive a timely response from the upstream server.

*/
// import (
// 	"encoding/json"
// 	"math/rand"
// 	"net/http"
// 	"strings"
// 	"sync"
// )

// type Car struct {
// 	Id    int64
// 	Name  string
// 	Model string
// 	Brand string
// 	Yeatr int64
// 	Price float64
// }

// // {
// // 	"name": "X7",
// // 	"model": "v6 inline",
// // 	"brand": "BMW",
// // 	"year": 2024 ,
// // 	"price": 20000000
// // }

// var Cars = make(map[int64]Car)
// var mu sync.Mutex

// func carHandeler(r *http.Request, w http.ResponseWriter) {

// 	path := r.URL.Path

// 	entity := strings.TrimPrefix(path, "/cars")
// 	entity = strings.Trim(entity, "/")

// 	switch r.Method {
// 	case "POST":
// 		if entity == "" {
// 			//create handelr
// 		} else {
// 			http.Error(w, "Incorect post Request", http.StatusBadRequest)
// 		}
// 	case "GET":
// 		if entity == "" {
// 			//create handeler
// 		} else {
// 			http.Error(w, " We dont Support this API ", http.StatusBadRequest)
// 		}
// 	case "PUT":
// 		if entity == "" {
// 			http.Error(w, " We dont Support this API ", http.StatusBadRequest)
// 		} else {
// 			//put handler
// 		}
// 	}
// }

// func createCar(r *http.Request, w http.ResponseWriter) {
// 	mu.Lock()

// 	defer mu.Unlock()

// 	var car Car
// 	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
// 		http.Error(w, "incorrect json input", http.StatusBadRequest)
// 	}

// 	id := rand.Intn(10000)
// 	car.Id = int64(id)
// 	//store in the map
// 	Cars[car.Id] = car

// 	w.Header().Set("Content-Type", "Application/json")
// 	w.WriteHeader((http.StatusAccepted))
// 	json.NewEncoder(w).Encode(car)
// }

// func getCar(r *http.Request, w http.ResponseWriter, id int64) {
// 	mu.Lock()

// 	defer mu.Unlock()

// 	car, ok := Cars[id]
// 	if !ok {
// 		http.Error(w, "Not found car", http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "Application/json")
// 	w.WriteHeader((http.StatusOK))
// 	json.NewEncoder(w).Encode(car)
// }
// func main() {

// }

type Car struct {
	ID    int
	Name  string
	Brand string
	Model string
	Year  int
	Price int
}

var Cars = make(map[int]Car)

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
			http.Error(w, "Incorrect request", http.StatusBadRequest)
		}
	case "GET":
		if entity == "" {
			http.Error(w, "Car is not found", http.StatusBadRequest)
		} else {
			id, _ := strconv.Atoi(entity)
			getCar(w, id)
		}
	case "DELETE":
		if entity == "" {
			http.Error(w, "Car is not found in the entity or deleted", http.StatusNoContent)
		} else {
			id, _ := strconv.Atoi(entity)
			deleteCar(w, id)
		}
	}
}

func createCar(w http.ResponseWriter, r *http.Request) {
	mu.Lock()

	defer mu.Unlock()

	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Incorrectnjson input ", http.StatusBadRequest)
		return
	}

	id := rand.Intn(1000)

	car.ID = id
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
		http.Error(w, "the requested car is not present ", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(car)
}

func deleteCar(w http.ResponseWriter, id int) {
	mu.Lock()

	defer mu.Unlock()

	_, ok := Cars[id]

	if !ok {
		http.Error(w, "Car is not present in the System", http.StatusBadRequest)
	}
	delete(Cars, id)
	w.Header().Set("Contnet-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/cars", carHandler)
	http.HandleFunc("/", carHandler)

	fmt.Println("Server Started to serve the API Request...")
	http.ListenAndServe(":3033", nil)
}
