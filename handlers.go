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
	customers, _ := routerHandler.service.GetAllCustomers()
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
		customer = domain.Customer{Id: "1", Name: "Krishna", City: "Gwalior", Zipcode: "474006", DateOfBirth: "31-08-1994", Status: "single"}
	} else {
		customer = domain.Customer{Id: "2", Name: "Madhur", City: "Morena", Zipcode: "474001", DateOfBirth: "21-11-1994", Status: "Single"}
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
