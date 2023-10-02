package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func envSanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variables are not defined")
	}
}

func Start() {

	envSanityCheck()
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
	router.HandleFunc("/customers/{customer_id:[0-9]+}", routerHandler.getCustomer) //by default the http method type would be GET here

	router.HandleFunc("/customers", routerHandler.createCustomer).Methods(http.MethodPost)
	//Here we are starting the servers

	server_address := os.Getenv("SERVER_ADDRESS")
	server_port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", server_address, server_port), router))
}
