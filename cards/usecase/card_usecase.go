package usecase

import (
	"github.com/kiranraj27/sprint-planner/internal/domain"
)

type cardUseCase struct {
	cardRepo domain.CardRepository
}

func NewCardUseCase(cardRepo domain.CardRepository) domain.CardUseCase {
	return &cardUseCase{cardRepo: cardRepo}
}

func (uc *cardUseCase) GetCards() ([]*domain.Card, error) {
	return uc.cardRepo.GetCards()
}
