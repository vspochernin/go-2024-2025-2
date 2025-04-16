package models

import (
	"time"
)

type Account struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AccountCreateRequest struct {
	UserID   int     `json:"user_id" validate:"required"`
	Currency string  `json:"currency" validate:"required,len=3"`
}

type AccountResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type AccountTransferRequest struct {
	FromAccountID int     `json:"from_account_id" validate:"required"`
	ToAccountID   int     `json:"to_account_id" validate:"required"`
	Amount        float64 `json:"amount" validate:"required,gt=0"`
}

func (a *AccountCreateRequest) Validate() error {
	if a.UserID <= 0 {
		return ErrInvalidUserID
	}
	if a.Currency != "RUB" { // По условию проекта поддерживаем только RUB
		return ErrInvalidCurrency
	}
	return nil
}

func (a *AccountTransferRequest) Validate() error {
	if a.FromAccountID <= 0 || a.ToAccountID <= 0 {
		return ErrInvalidAccountID
	}
	if !ValidateAmount(a.Amount) {
		return ErrInvalidAmount
	}
	return nil
}

func (a *Account) ToResponse() AccountResponse {
	return AccountResponse{
		ID:        a.ID,
		UserID:    a.UserID,
		Balance:   a.Balance,
		Currency:  a.Currency,
		CreatedAt: a.CreatedAt,
	}
} 