package restapimiddilewire

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/cars", carHandler)
	http.HandleFunc("/cars/", carHandler)

	fmt.Println("Serever started running and listen")
	http.ListenAndServe(":3015", nil)
}
