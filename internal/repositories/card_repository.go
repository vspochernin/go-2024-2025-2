package repositories

import (
	"banksystem/internal/models"
	"database/sql"
)

type CardRepository struct {
	db *sql.DB
}

func NewCardRepository(db *sql.DB) *CardRepository {
	return &CardRepository{db: db}
}

func (r *CardRepository) Create(card *models.Card) error {
	query := `
		INSERT INTO cards (account_id, encrypted_data, hashed_cvv, hmac, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	return r.db.QueryRow(
		query,
		card.AccountID,
		card.EncryptedData,
		card.HashedCVV,
		card.HMAC,
		card.CreatedAt,
	).Scan(&card.ID)
}

func (r *CardRepository) GetByID(id int64) (*models.Card, error) {
	card := &models.Card{}
	
	query := `
		SELECT id, account_id, encrypted_data, hashed_cvv, hmac, created_at
		FROM cards
		WHERE id = $1
	`
	
	err := r.db.QueryRow(query, id).Scan(
		&card.ID,
		&card.AccountID,
		&card.EncryptedData,
		&card.HashedCVV,
		&card.HMAC,
		&card.CreatedAt,
	)
	
	if err == sql.ErrNoRows {
		return nil, nil
	}
	
	if err != nil {
		return nil, err
	}
	
	return card, nil
}

func (r *CardRepository) GetByUserID(userID int64) ([]*models.Card, error) {
	query := `
		SELECT c.id, c.account_id, c.encrypted_data, c.hashed_cvv, c.hmac, c.created_at
		FROM cards c
		JOIN accounts a ON c.account_id = a.id
		WHERE a.user_id = $1
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []*models.Card
	for rows.Next() {
		card := &models.Card{}
		err := rows.Scan(
			&card.ID,
			&card.AccountID,
			&card.EncryptedData,
			&card.HashedCVV,
			&card.HMAC,
			&card.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

func (r *CardRepository) GetByAccountID(accountID int64) ([]*models.Card, error) {
	query := `
		SELECT id, account_id, encrypted_data, hashed_cvv, hmac, created_at
		FROM cards
		WHERE account_id = $1
	`

	rows, err := r.db.Query(query, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []*models.Card
	for rows.Next() {
		card := &models.Card{}
		err := rows.Scan(
			&card.ID,
			&card.AccountID,
			&card.EncryptedData,
			&card.HashedCVV,
			&card.HMAC,
			&card.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	return cards, nil
}

func (r *CardRepository) VerifyHMAC(id int64, hmac []byte) (bool, error) {
	query := `
		SELECT hmac = $1
		FROM cards
		WHERE id = $2
	`

	var matches bool
	err := r.db.QueryRow(query, hmac, id).Scan(&matches)
	if err != nil {
		return false, err
	}

	return matches, nil
}

func (r *CardRepository) GetByAccountUserID(userID int64) ([]*models.Card, error) {
	query := `
		SELECT c.id, c.account_id, c.encrypted_data, c.hashed_cvv, c.hmac, c.created_at
		FROM cards c
		JOIN accounts a ON c.account_id = a.id
		WHERE a.user_id = $1
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []*models.Card
	for rows.Next() {
		card := &models.Card{}
		err := rows.Scan(
			&card.ID,
			&card.AccountID,
			&card.EncryptedData,
			&card.HashedCVV,
			&card.HMAC,
			&card.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	return cards, nil
} 