package main

import (
	"fmt"
	"net/http"
)

func app(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " Hey🫡 Welcome to myswiggyFoodDeliveryApplicaion")
}

func main() {
	http.HandleFunc("/", app)
	http.ListenAndServe(":8080", nil)
}
