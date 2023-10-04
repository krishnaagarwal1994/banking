package domain

import "banking/dto"

type Transaction struct {
	TransactionID   string
	AccountID       int64
	Amount          float64
	TransactionType string
	TransactionDate string
}

func (t Transaction) ToTransactionResponseDTO() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionID: t.TransactionID,
	}
}
