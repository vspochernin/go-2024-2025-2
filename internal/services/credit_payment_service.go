package services

import (
	"banksystem/internal/models"
	"banksystem/internal/repositories"
	"time"
)

type CreditPaymentService struct {
	paymentRepo *repositories.CreditPaymentRepository
	creditRepo  *repositories.CreditRepository
	accountRepo *repositories.AccountRepository
}

func NewCreditPaymentService(
	paymentRepo *repositories.CreditPaymentRepository,
	creditRepo *repositories.CreditRepository,
	accountRepo *repositories.AccountRepository,
) *CreditPaymentService {
	return &CreditPaymentService{
		paymentRepo: paymentRepo,
		creditRepo:  creditRepo,
		accountRepo: accountRepo,
	}
}

func (s *CreditPaymentService) CreatePayment(creditID int, amount float64, paymentDate time.Time) error {
	payment := &models.CreditPayment{
		CreditID:    creditID,
		Amount:      amount,
		PaymentDate: paymentDate,
		Status:      "pending",
		CreatedAt:   time.Now(),
	}

	return s.paymentRepo.Create(payment)
}

func (s *CreditPaymentService) ProcessPayment(paymentID int) error {
	payment, err := s.paymentRepo.GetByID(paymentID)
	if err != nil {
		return err
	}

	credit, err := s.creditRepo.GetByID(payment.CreditID)
	if err != nil {
		return err
	}

	account, err := s.accountRepo.GetByID(credit.AccountID)
	if err != nil {
		return err
	}

	if account.Balance < payment.Amount {
		return s.paymentRepo.UpdateStatus(paymentID, "failed")
	}

	// Списание средств
	account.Balance -= payment.Amount
	if err := s.accountRepo.Update(account); err != nil {
		return err
	}

	// Обновление статуса платежа
	return s.paymentRepo.UpdateStatus(paymentID, "completed")
}

func (s *CreditPaymentService) GetPaymentsByCreditID(creditID int) ([]*models.CreditPayment, error) {
	return s.paymentRepo.GetByCreditID(creditID)
}

func (s *CreditPaymentService) GetPendingPayments() ([]*models.CreditPayment, error) {
	return s.paymentRepo.GetPendingPayments()
} 