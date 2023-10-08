package tests

import (
	"banking/dto"
	"net/http"
	"testing"
)

func TestShouldReturnErrorWhenTransactionTypeNotWithdrawalOrDeposit(t *testing.T) {
	// Arrange
	request := dto.TransactionRequest{
		TransactionType: "unknown",
	}
	// Act
	validationError := request.Validate()

	// Assertion
	if validationError.Code != http.StatusUnprocessableEntity {
		t.Error("Incorrect error code")
	}

	if validationError.Message != "Invalid transaction type, only withdrawal and deposit transaction type are allowed." {
		t.Error("Incorrect error message")
	}
}
