package dto

import (
	"banking/errs"
	"strings"
)

type TransactionRequest struct {
	AccountID       int64   `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func (t TransactionRequest) Validate() *errs.AppError {
	transactionType := strings.ToLower(t.TransactionType)
	if transactionType != "withdrawal" && transactionType != "deposit" {
		return errs.NewUnexpectedError("Invalid transaction type, only withdrawal and deposit transaction type are allowed.")
	}
	if t.Amount < 0 {
		return errs.NewUnexpectedError("Amount value can't be less than zero")
	}
	return nil
}
