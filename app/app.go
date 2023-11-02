package app

import (
	"banking/domain"
	"banking/service"
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var vp *viper.Viper = viper.New()

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
}

func loadConfiguration() (Config, error) {
	vp.SetConfigName("configs")
	vp.AddConfigPath(".")
	vp.SetConfigType("env")

	err := vp.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
		return Config{}, err
	}
	var config Config
	err = vp.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
		return Config{}, err
	}
	return config, nil
}

func Start() {

	config, err := loadConfiguration()

	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
		panic(err)
	}

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
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", config.ServerAddress, config.ServerPort), router))
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
