package models

import "time"

type CreditPayment struct {
	ID          int       `json:"id"`
	CreditID    int       `json:"credit_id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
} 