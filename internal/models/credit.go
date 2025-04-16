package models

import (
	"time"
)

type Credit struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	AccountID    int       `json:"account_id"`
	Amount       float64   `json:"amount"`
	InterestRate float64   `json:"interest_rate"`
	TermMonths   int       `json:"term_months"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreditCreateRequest struct {
	UserID       int     `json:"user_id" validate:"required"`
	AccountID    int     `json:"account_id" validate:"required"`
	Amount       float64 `json:"amount" validate:"required,gt=0"`
	TermMonths   int     `json:"term_months" validate:"required,gt=0"`
}

type CreditResponse struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	AccountID    int       `json:"account_id"`
	Amount       float64   `json:"amount"`
	InterestRate float64   `json:"interest_rate"`
	TermMonths   int       `json:"term_months"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

const (
	CreditStatusActive    = "ACTIVE"
	CreditStatusClosed    = "CLOSED"
	CreditStatusOverdue   = "OVERDUE"
	CreditStatusRejected  = "REJECTED"
)

func (c *CreditCreateRequest) Validate() error {
	if c.UserID <= 0 {
		return ErrInvalidUserID
	}
	if c.AccountID <= 0 {
		return ErrInvalidAccountID
	}
	if !ValidateAmount(c.Amount) {
		return ErrInvalidAmount
	}
	if c.TermMonths <= 0 {
		return ErrInvalidTerm
	}
	return nil
}

func (c *Credit) ToResponse() CreditResponse {
	return CreditResponse{
		ID:           c.ID,
		UserID:       c.UserID,
		AccountID:    c.AccountID,
		Amount:       c.Amount,
		InterestRate: c.InterestRate,
		TermMonths:   c.TermMonths,
		Status:       c.Status,
		CreatedAt:    c.CreatedAt,
	}
}

func ValidateCreditStatus(status string) bool {
	switch status {
	case CreditStatusActive,
		CreditStatusClosed,
		CreditStatusOverdue,
		CreditStatusRejected:
		return true
	default:
		return false
	}
} 