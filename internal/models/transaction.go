package models

import (
	"time"
)

type Transaction struct {
	ID            int       `json:"id"`
	FromAccountID int       `json:"from_account_id"`
	ToAccountID   int       `json:"to_account_id"`
	Amount        float64   `json:"amount"`
	Type          string    `json:"type"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

type TransactionResponse struct {
	ID            int       `json:"id"`
	FromAccountID int       `json:"from_account_id"`
	ToAccountID   int       `json:"to_account_id"`
	Amount        float64   `json:"amount"`
	Type          string    `json:"type"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

const (
	TransactionTypeDeposit    = "DEPOSIT"
	TransactionTypeWithdrawal = "WITHDRAWAL"
	TransactionTypeTransfer   = "TRANSFER"
	TransactionTypePayment    = "PAYMENT"

	TransactionStatusPending   = "PENDING"
	TransactionStatusCompleted = "COMPLETED"
	TransactionStatusFailed    = "FAILED"
)

func (t *Transaction) ToResponse() TransactionResponse {
	return TransactionResponse{
		ID:            t.ID,
		FromAccountID: t.FromAccountID,
		ToAccountID:   t.ToAccountID,
		Amount:        t.Amount,
		Type:          t.Type,
		Status:        t.Status,
		CreatedAt:     t.CreatedAt,
	}
}

func ValidateTransactionType(tType string) bool {
	switch tType {
	case TransactionTypeDeposit,
		TransactionTypeWithdrawal,
		TransactionTypeTransfer,
		TransactionTypePayment:
		return true
	default:
		return false
	}
}

func ValidateTransactionStatus(status string) bool {
	switch status {
	case TransactionStatusPending,
		TransactionStatusCompleted,
		TransactionStatusFailed:
		return true
	default:
		return false
	}
} 