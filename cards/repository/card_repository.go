package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/kiranraj27/sprint-planner/internal/domain"
)

type cardRepository struct {
	db *sqlx.DB
}

func NewCardRepository(db *sqlx.DB) domain.CardRepository {
	return &cardRepository{db: db}
}

func (r *cardRepository) GetCards() ([]*domain.Card, error) {
	var cards []*domain.Card
	query := "SELECT id, value, set_type FROM cards"
	err := r.db.Select(&cards, query)
	return cards, err
}
