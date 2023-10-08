package app

import (
	"banking/domain"
	"banking/service"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
	sqlClient := getSQLClient()

	customerRepositoryDb := domain.NewCustomerRepositoryDb(sqlClient)
	accountRepositoryDB := domain.NewAccountRepositoryDB(sqlClient)

	customerService := service.NewCustomerService(customerRepositoryDb)
	accountService := service.NewAccountService(accountRepositoryDB)

	customerHandler := RouterHandler{service: customerService}
	accountHandler := AccountHandler{accountService: accountService}

	router.HandleFunc("/greet", greet)

	//Registering an enpoint to fetch all customers
	router.HandleFunc("/customers", customerHandler.getAllCustomer).Methods(http.MethodGet).
		Name("GetAllCustomers")

	//Registering an endpoint to return customer based on customer id
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.getCustomer).
		Name("GetCustomer") //by default the http method type would be GET here

	//Registering an endpoint to create a new bank account for the given customer
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", accountHandler.createNewAccount).
		Methods(http.MethodPost).Name("NewAccount")

	//Registering an endpoint to create a new bank account for the account id passed in the request payload.
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", accountHandler.NewTransaction).
		Methods(http.MethodPost).Name("NewTransaction")

	//Injecting the middleware for performing the authorisation validation
	authRepository := domain.NewAuthRepository()
	authMiddleware := AuthMiddleware{repository: authRepository}
	router.Use(authMiddleware.authorizationHandler())

	//Here we are starting the servers
	server_address := os.Getenv("SERVER_ADDRESS")
	server_port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", server_address, server_port), router))
}

func getSQLClient() *sql.DB {
	db, err := sql.Open("mysql", "root:Gn1d0c@123@/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
