package app

import (
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
	customer, err := routerHandler.service.GetCustomerWith(customer_id)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, t interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(t)
	if err != nil {
		panic(err)
	}
}
