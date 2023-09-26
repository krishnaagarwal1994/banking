package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Here we are defining the routes
	http.HandleFunc("/greet", greet)

	//Registering an enpoint to fetch all customers
	http.HandleFunc("/customers", getAllCustomer)

	//Here we are starting the servers
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Krishna", "Gwalior", "474006"},
		{"Madhur", "Morena", "474001"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
