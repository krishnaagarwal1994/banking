package dto

import (
	"banking/errs"
	"strings"
)

type NewAccountRequest struct {
	AccountType string  `json:"account_type"`
	CustomerID  string  `json:"customer_id"`
	Amount      float64 `json:"amount"`
}

func (requestDTO NewAccountRequest) Validate() *errs.AppError {
	if strings.ToLower(requestDTO.AccountType) != "savings" && strings.ToLower(requestDTO.AccountType) != "checking" {
		return errs.NewUnexpectedError("account should be either savings or checking")
	}
	if requestDTO.Amount < 5000 {
		return errs.NewUnexpectedError("Minimum deposit amount to open a new bank account is 5000.00")
	}
	return nil
}
