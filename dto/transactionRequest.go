package dto

import (
	"banking/errs"
	"strings"
)

type TransactionRequest struct {
	AccountID       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerID      string  `json:"-"`
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

func (t TransactionRequest) IsTransactionTypeWithdrawal() bool {
	if strings.ToLower(t.TransactionType) == "withdrawal" {
		return true
	} else {
		return false
	}
}
