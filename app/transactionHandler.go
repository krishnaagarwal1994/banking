package app

import (
	"banking/dto"
	"banking/logger"
	"banking/service"
	"encoding/json"
	"net/http"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func (handler TransactionHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {
	var transactionRequestDTO dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&transactionRequestDTO)
	if err != nil {
		logger.Error("Unable to map the request body to TransactionResponseDTO")
		writeResponse(w, http.StatusBadRequest, err.Error())
	}
	transactionResponse, error := handler.transactionService.NewTransaction(transactionRequestDTO)
	if error != nil {
		writeResponse(w, error.Code, error.AsMessage())
	} else {
		writeResponse(w, http.StatusCreated, transactionResponse)
	}
}
