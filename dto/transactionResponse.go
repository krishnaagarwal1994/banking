package dto

type TransactionResponse struct {
	UpdatedBalance  float64 `json:"updated_balance,omitempty"`
	TransactionID   string  `json:"transaction_id,omitempty"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	AccountID       string  `json:"account_id"`
}
