package repositories

import (
	"banksystem/internal/models"
	"database/sql"
	"time"
)

type AccountRepository struct {
	*BaseRepository
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *AccountRepository) Create(account *models.Account) error {
	query := `
		INSERT INTO accounts (user_id, balance, currency, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		account.UserID,
		account.Balance,
		account.Currency,
		time.Now(),
		time.Now(),
	).Scan(&account.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) GetByID(id int) (*models.Account, error) {
	query := `
		SELECT id, user_id, balance, currency, created_at, updated_at
		FROM accounts
		WHERE id = $1
	`

	account := &models.Account{}
	err := r.db.QueryRow(query, id).Scan(
		&account.ID,
		&account.UserID,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (r *AccountRepository) GetByUserID(userID int) ([]*models.Account, error) {
	query := `
		SELECT id, user_id, balance, currency, created_at, updated_at
		FROM accounts
		WHERE user_id = $1
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*models.Account
	for rows.Next() {
		account := &models.Account{}
		err := rows.Scan(
			&account.ID,
			&account.UserID,
			&account.Balance,
			&account.Currency,
			&account.CreatedAt,
			&account.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (r *AccountRepository) UpdateBalance(tx *sql.Tx, accountID int, amount float64) error {
	query := `
		UPDATE accounts
		SET balance = balance + $1,
			updated_at = $2
		WHERE id = $3
		RETURNING balance
	`

	var balance float64
	err := tx.QueryRow(query, amount, time.Now(), accountID).Scan(&balance)
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) Transfer(fromAccountID, toAccountID int, amount float64) error {
	tx, err := r.BeginTx()
	if err != nil {
		return err
	}
	defer r.RollbackTx(tx)

	// Проверяем достаточность средств
	var fromBalance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = $1", fromAccountID).Scan(&fromBalance)
	if err != nil {
		return err
	}
	if fromBalance < amount {
		return ErrInvalidData
	}

	// Списание средств
	if err := r.UpdateBalance(tx, fromAccountID, -amount); err != nil {
		return err
	}

	// Зачисление средств
	if err := r.UpdateBalance(tx, toAccountID, amount); err != nil {
		return err
	}

	return r.CommitTx(tx)
}

func (r *AccountRepository) Update(account *models.Account) error {
	query := `
		UPDATE accounts
		SET balance = $1,
			updated_at = $2
		WHERE id = $3
	`

	_, err := r.db.Exec(
		query,
		account.Balance,
		time.Now(),
		account.ID,
	)

	if err != nil {
		return err
	}

	return nil
} 