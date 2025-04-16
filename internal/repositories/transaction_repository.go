package repositories

import (
	"banksystem/internal/models"
	"database/sql"
	"time"
)

type TransactionRepository struct {
	*BaseRepository
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	query := `
		INSERT INTO transactions (from_account_id, to_account_id, amount, type, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		transaction.FromAccountID,
		transaction.ToAccountID,
		transaction.Amount,
		transaction.Type,
		transaction.Status,
		time.Now(),
	).Scan(&transaction.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepository) GetByID(id int) (*models.Transaction, error) {
	query := `
		SELECT id, from_account_id, to_account_id, amount, type, status, created_at
		FROM transactions
		WHERE id = $1
	`

	transaction := &models.Transaction{}
	err := r.db.QueryRow(query, id).Scan(
		&transaction.ID,
		&transaction.FromAccountID,
		&transaction.ToAccountID,
		&transaction.Amount,
		&transaction.Type,
		&transaction.Status,
		&transaction.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *TransactionRepository) GetByAccountID(accountID int) ([]*models.Transaction, error) {
	query := `
		SELECT id, from_account_id, to_account_id, amount, type, status, created_at
		FROM transactions
		WHERE from_account_id = $1 OR to_account_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		transaction := &models.Transaction{}
		err := rows.Scan(
			&transaction.ID,
			&transaction.FromAccountID,
			&transaction.ToAccountID,
			&transaction.Amount,
			&transaction.Type,
			&transaction.Status,
			&transaction.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *TransactionRepository) UpdateStatus(id int, status string) error {
	query := `
		UPDATE transactions
		SET status = $1
		WHERE id = $2
	`

	_, err := r.db.Exec(query, status, id)
	if err != nil {
		return err
	}

	return nil
} 