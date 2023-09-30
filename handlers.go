package main

import (
	"banking/domain"
	"banking/service"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type RouterHandler struct {
	service service.CustomerService
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func (routerHandler *RouterHandler) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	// return routerHandler.service.FindAll()
	customers := []domain.Customer{
		{Id: "1", Name: "Krishna", City: "Gwalior", Zipcode: "474006"},
		{Id: "2", Name: "Madhur", City: "Morena", Zipcode: "474001"},
	}

	contentType := r.Header.Get("Content-Type")
	if contentType == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		// Logic to add the custom headers to the response.
		w.Header().Add("Content-Type", "application/json")
		// Logic to convert the models to JSON
		json.NewEncoder(w).Encode(customers)
	}
}

func (routerHandler *RouterHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	customer_id := requestVars["customer_id"]
	fmt.Print(customer_id)
	var customer domain.Customer
	if customer_id == "1" {
		customer = domain.Customer{Id: "1", Name: "Krishna", City: "Gwalior", Zipcode: "474006"}
	} else {
		customer = domain.Customer{Id: "2", Name: "Madhur", City: "Morena", Zipcode: "474001"}
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func (routerHandler *RouterHandler) createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
	//Sample business logic to decode the request payload
	// func test(rw http.ResponseWriter, req *http.Request) {
	// 	decoder := json.NewDecoder(req.Body)
	// 	var t test_struct
	// 	err := decoder.Decode(&t)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	log.Println(t.Test)
	// }
}
