package main

import (
	"banking/domain"
	"banking/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func registerHandlers() {
	router := mux.NewRouter()

	// This is the stubbed repository which would give hard coded response
	// customerRepositoryStub := domain.NewCustomerRepositoryStub()

	// This is the Actual Adapter and would return the data from the DB
	customerRepositoryDb := domain.NewCustomerRepositoryDb()

	customerService := service.NewCustomerService(customerRepositoryDb)

	routerHandler := RouterHandler{service: customerService}

	router.HandleFunc("/greet", greet)

	//Registering an enpoint to fetch all customers
	router.HandleFunc("/customers", routerHandler.getAllCustomer).Methods(http.MethodGet)

	//Registering an endpoint to return customer based on customer id
	router.HandleFunc("/customers/{customer_id}", routerHandler.getCustomer) //by default the http method type would be GET here

	router.HandleFunc("/customers", routerHandler.createCustomer).Methods(http.MethodPost)
	//Here we are starting the servers
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
