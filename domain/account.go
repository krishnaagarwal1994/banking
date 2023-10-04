package domain

import (
	"banking/dto"
)

type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (account Account) ToNewAccountResponseDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountID: account.AccountID}
}

func (account Account) ToFetchAccountDetailsResponseDTO() dto.FetchAccountDetailsResponse {
	return dto.FetchAccountDetailsResponse{
		AccountID:   account.AccountID,
		CustomerID:  account.CustomerID,
		OpeningDate: account.OpeningDate,
		AccountType: account.AccountType,
		Amount:      account.Amount,
		Status:      account.Status,
	}
}

func (account Account) CanWithDraw(amount int) bool {
	return account.Amount >= float64(amount)
}
