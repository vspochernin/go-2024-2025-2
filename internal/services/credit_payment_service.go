package services

import (
	"banksystem/internal/models"
	"banksystem/internal/repositories"
	"context"
	"database/sql"
	"time"
)

type CreditPaymentService struct {
	paymentRepo *repositories.CreditPaymentRepository
	creditRepo  *repositories.CreditRepository
	accountRepo *repositories.AccountRepository
	db         *sql.DB
}

func NewCreditPaymentService(
	db *sql.DB,
	paymentRepo *repositories.CreditPaymentRepository,
	creditRepo *repositories.CreditRepository,
	accountRepo *repositories.AccountRepository,
) *CreditPaymentService {
	return &CreditPaymentService{
		db:          db,
		paymentRepo: paymentRepo,
		creditRepo:  creditRepo,
		accountRepo: accountRepo,
	}
}

func (s *CreditPaymentService) CreatePayment(ctx context.Context, creditID int64, amount float64, dueDate time.Time) (*models.CreditPayment, error) {
	payment := &models.CreditPayment{
		CreditID: creditID,
		Amount:   amount,
		Status:   "pending",
		DueDate:  dueDate,
	}

	err := s.paymentRepo.Create(ctx, payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *CreditPaymentService) ProcessPayment(ctx context.Context, paymentID int64) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	payment, err := s.paymentRepo.GetByID(ctx, paymentID)
	if err != nil {
		return err
	}

	if payment.Status != "pending" {
		return nil
	}

	credit, err := s.creditRepo.GetByID(int(payment.CreditID))
	if err != nil {
		return err
	}

	account, err := s.accountRepo.GetByID(ctx, credit.AccountID)
	if err != nil {
		return err
	}

	if account.Balance < payment.Amount {
		return s.paymentRepo.UpdateStatus(ctx, paymentID, "failed")
	}

	// Списываем средства со счета
	account.Balance -= payment.Amount
	err = s.accountRepo.Update(ctx, tx, account)
	if err != nil {
		return err
	}

	err = s.paymentRepo.UpdateStatus(ctx, paymentID, "completed")
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *CreditPaymentService) GetPaymentsByCreditID(ctx context.Context, creditID int64) ([]*models.CreditPayment, error) {
	return s.paymentRepo.GetByCreditID(ctx, creditID)
}

func (s *CreditPaymentService) GetPendingPayments(ctx context.Context) ([]*models.CreditPayment, error) {
	return s.paymentRepo.GetPending(ctx)
}

func (s *CreditPaymentService) UpdateStatus(ctx context.Context, paymentID int64, status string) error {
	return s.paymentRepo.UpdateStatus(ctx, paymentID, status)
} 