package service

import (
	"banking/dto"
	"banking/errs"
)

type AccountService interface {
	NewAccount(requestDTO dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}
