package repositories

import (
	"banksystem/internal/models"
	"database/sql"
	"time"
)

type CreditPaymentRepository struct {
	*BaseRepository
}

func NewCreditPaymentRepository(db *sql.DB) *CreditPaymentRepository {
	return &CreditPaymentRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *CreditPaymentRepository) Create(payment *models.CreditPayment) error {
	query := `
		INSERT INTO credit_payments (credit_id, amount, payment_date, status, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		payment.CreditID,
		payment.Amount,
		payment.PaymentDate,
		payment.Status,
		time.Now(),
	).Scan(&payment.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *CreditPaymentRepository) GetByID(id int) (*models.CreditPayment, error) {
	query := `
		SELECT id, credit_id, amount, payment_date, status, created_at
		FROM credit_payments
		WHERE id = $1
	`

	payment := &models.CreditPayment{}
	err := r.db.QueryRow(query, id).Scan(
		&payment.ID,
		&payment.CreditID,
		&payment.Amount,
		&payment.PaymentDate,
		&payment.Status,
		&payment.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (r *CreditPaymentRepository) GetByCreditID(creditID int) ([]*models.CreditPayment, error) {
	query := `
		SELECT id, credit_id, amount, payment_date, status, created_at
		FROM credit_payments
		WHERE credit_id = $1
		ORDER BY payment_date ASC
	`

	rows, err := r.db.Query(query, creditID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*models.CreditPayment
	for rows.Next() {
		payment := &models.CreditPayment{}
		err := rows.Scan(
			&payment.ID,
			&payment.CreditID,
			&payment.Amount,
			&payment.PaymentDate,
			&payment.Status,
			&payment.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

func (r *CreditPaymentRepository) UpdateStatus(id int, status string) error {
	query := `
		UPDATE credit_payments
		SET status = $1
		WHERE id = $2
	`

	_, err := r.db.Exec(query, status, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *CreditPaymentRepository) GetPendingPayments() ([]*models.CreditPayment, error) {
	query := `
		SELECT id, credit_id, amount, payment_date, status, created_at
		FROM credit_payments
		WHERE status = 'pending'
		AND payment_date <= $1
		ORDER BY payment_date ASC
	`

	rows, err := r.db.Query(query, time.Now())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*models.CreditPayment
	for rows.Next() {
		payment := &models.CreditPayment{}
		err := rows.Scan(
			&payment.ID,
			&payment.CreditID,
			&payment.Amount,
			&payment.PaymentDate,
			&payment.Status,
			&payment.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
} 