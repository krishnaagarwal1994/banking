package app

import (
	"banking/dto"
	"banking/logger"
	"banking/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	accountService service.AccountService
}

func (handler AccountHandler) createNewAccount(w http.ResponseWriter, r *http.Request) {
	// Decoding the request to our Request DTO
	var requestDTO dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&requestDTO)
	requestDTO.CustomerID = getCustomerIDFromRequest(r)
	if err != nil {
		logger.Error("Unable to decode the request body to request DTO")
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		newAccountDTOResponse, err := handler.accountService.NewAccount(requestDTO)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, newAccountDTOResponse)
		}
	}
}

func (handler AccountHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {
	requestVars := mux.Vars(r)
	accountID := requestVars["account_id"]
	customerID := requestVars["customer_id"]
	var transactionRequestDTO dto.TransactionRequest
	//Mapping the request body to TransactionRequestDTO
	err := json.NewDecoder(r.Body).Decode(&transactionRequestDTO)
	if err != nil {
		logger.Error("Unable to map the request body to TransactionResponseDTO")
		writeResponse(w, http.StatusBadRequest, err.Error())
	}
	transactionRequestDTO.CustomerID = customerID
	transactionRequestDTO.AccountID = accountID
	transactionResponse, error := handler.accountService.NewTransaction(transactionRequestDTO)
	if error != nil {
		writeResponse(w, error.Code, error.AsMessage())
	} else {
		writeResponse(w, http.StatusCreated, transactionResponse)
	}
}

func getCustomerIDFromRequest(r *http.Request) string {
	requestParams := mux.Vars(r)
	print(requestParams)
	customerId := requestParams["customer_id"]
	return customerId
}
