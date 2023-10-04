package domain

import (
	"banking/dto"
	"strings"
)

type Transaction struct {
	TransactionID   string
	AccountID       string
	Amount          float64
	TransactionType string
	TransactionDate string
}

func (t Transaction) ToTransactionResponseDTO() dto.TransactionResponse {
	return dto.TransactionResponse{
		UpdatedBalance:  t.Amount,
		TransactionID:   t.TransactionID,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
		AccountID:       t.AccountID,
	}
}

func (t Transaction) IsWithdrawal() bool {
	return strings.ToLower(t.TransactionType) == "withdrawal"
}
