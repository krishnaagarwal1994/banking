package service

import (
	"banking/dto"
	"banking/errs"
)

type TransactionService interface {
	NewTransaction(r dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}
