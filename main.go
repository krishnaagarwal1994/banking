package main

import (
	"fmt"
	"log"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Krishna", "Gwalior", "474006"},
		{"Madhur", "Morena", "474001"},
	}
}

func main() {
	// Here we are defining the routes
	http.HandleFunc("/greet", greet)

	//Here we are starting the servers
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
