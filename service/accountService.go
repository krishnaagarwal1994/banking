package service

import (
	"banking/dto"
	"banking/errs"
)

type AccountService interface {
	NewAccount(requestDTO dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	AccountDetails(requestDTO dto.FetchAccountDetailsRequest) (*dto.FetchAccountDetailsResponse, *errs.AppError)
	NewTransaction(r dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}
