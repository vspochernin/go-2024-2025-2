package models

import (
	"time"
)

type Card struct {
	ID         int       `json:"id"`
	AccountID  int       `json:"account_id"`
	CardNumber []byte    `json:"-"` // Зашифрованный номер карты
	ExpiryDate []byte    `json:"-"` // Зашифрованная дата
	CVVHash    string    `json:"-"` // Хеш CVV
	HMAC       []byte    `json:"-"` // HMAC для проверки целостности
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CardCreateRequest struct {
	AccountID  int    `json:"account_id" validate:"required"`
	CardNumber string `json:"card_number" validate:"required"`
	ExpiryDate string `json:"expiry_date" validate:"required"`
	CVV        string `json:"cvv" validate:"required"`
}

type CardResponse struct {
	ID         int       `json:"id"`
	AccountID  int       `json:"account_id"`
	CardNumber string    `json:"card_number"` // Расшифрованный номер карты
	ExpiryDate string    `json:"expiry_date"` // Расшифрованная дата
	CreatedAt  time.Time `json:"created_at"`
}

func (c *CardCreateRequest) Validate() error {
	if c.AccountID <= 0 {
		return ErrInvalidAccountID
	}
	if !ValidateCardNumber(c.CardNumber) {
		return ErrInvalidCardNumber
	}
	if !ValidateExpiryDate(c.ExpiryDate) {
		return ErrInvalidExpiryDate
	}
	if !ValidateCVV(c.CVV) {
		return ErrInvalidCVV
	}
	return nil
}

func (c *Card) ToResponse(cardNumber, expiryDate string) CardResponse {
	return CardResponse{
		ID:         c.ID,
		AccountID:  c.AccountID,
		CardNumber: cardNumber,
		ExpiryDate: expiryDate,
		CreatedAt:  c.CreatedAt,
	}
} 