package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"banking/logger"
	"time"
)

type DefaultAccountService struct {
	repository domain.AccountRepository
}

func (service DefaultAccountService) NewAccount(requestDTO dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	//First we will validate whether the recieved request meetings the criteria or not
	err := requestDTO.Validate()
	if err != nil {
		logger.Error("Request DTO validation failed")
		return nil, err
	}
	account := domain.Account{
		AccountID:   "",
		CustomerID:  requestDTO.CustomerID,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: requestDTO.AccountType,
		Amount:      requestDTO.Amount,
		Status:      "1",
	}
	newAccount, err := service.repository.Save(account)
	if err != nil {
		return nil, err
	}
	newAccountResponseDTO := newAccount.ToNewAccountResponseDTO()
	return &newAccountResponseDTO, nil
}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repository: repository}
}
