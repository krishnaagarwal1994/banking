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

func (service DefaultAccountService) AccountDetails(requestDTO dto.FetchAccountDetailsRequest) (*dto.FetchAccountDetailsResponse, *errs.AppError) {
	account, err := service.repository.FetchDetails(requestDTO.AccountID)
	if err != nil {
		return nil, err
	}
	accountDetailResponse := account.ToFetchAccountDetailsResponseDTO()
	return &accountDetailResponse, nil
}

func (service DefaultAccountService) NewTransaction(r dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	//Request validation
	validationError := r.Validate()
	if validationError != nil {
		logger.Error("Request validation failed")
		return nil, validationError
	}
	if r.IsTransactionTypeWithdrawal() {
		account, err := service.repository.FetchDetails(r.AccountID)
		if err != nil {
			return nil, err
		}
		if !account.CanWithDraw(int(r.Amount)) {
			return nil, errs.NewUnexpectedError("Insufficient Bank Balance")
		}
	}
	transaction := domain.Transaction{
		AccountID:       r.AccountID,
		Amount:          r.Amount,
		TransactionType: r.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	newTransaction, err := service.repository.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}
	transactionResponseDTO := newTransaction.ToTransactionResponseDTO()
	return &transactionResponseDTO, nil
}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repository: repository}
}
