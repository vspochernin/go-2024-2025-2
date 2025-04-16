package repositories

import (
	"banksystem/internal/models"
	"database/sql"
	"time"
)

type CardRepository struct {
	*BaseRepository
}

func NewCardRepository(db *sql.DB) *CardRepository {
	return &CardRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *CardRepository) Create(card *models.Card) error {
	query := `
		INSERT INTO cards (account_id, card_number, expiry_date, cvv_hash, hmac, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		card.AccountID,
		card.CardNumber,
		card.ExpiryDate,
		card.CVVHash,
		card.HMAC,
		time.Now(),
		time.Now(),
	).Scan(&card.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *CardRepository) GetByID(id int) (*models.Card, error) {
	query := `
		SELECT id, account_id, card_number, expiry_date, cvv_hash, hmac, created_at, updated_at
		FROM cards
		WHERE id = $1
	`

	card := &models.Card{}
	err := r.db.QueryRow(query, id).Scan(
		&card.ID,
		&card.AccountID,
		&card.CardNumber,
		&card.ExpiryDate,
		&card.CVVHash,
		&card.HMAC,
		&card.CreatedAt,
		&card.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return card, nil
}

func (r *CardRepository) GetByAccountID(accountID int) ([]*models.Card, error) {
	query := `
		SELECT id, account_id, card_number, expiry_date, cvv_hash, hmac, created_at, updated_at
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
			&card.CardNumber,
			&card.ExpiryDate,
			&card.CVVHash,
			&card.HMAC,
			&card.CreatedAt,
			&card.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	return cards, nil
}

func (r *CardRepository) VerifyHMAC(id int, hmac []byte) (bool, error) {
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