package dto

type TransactionResponse struct {
	UpdatedBalance float64 `json:"updated_balance,omitempty"`
	TransactionID  string  `json:"transaction_id,omitempty"`
}
