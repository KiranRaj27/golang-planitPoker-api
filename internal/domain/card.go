package domain

// Card represents a card with its value and set type.
// @Description Card entity containing value and set type.
// @Example {
//   "id": 1,
//   "value": "1",
//   "set_type": "scrum"
// }
type Card struct {
	ID     int    `db:"id" json:"id"`
	Value  string `db:"value" json:"value"`
	SetType string `db:"set_type" json:"set_type"`
}

type CardRepository interface {
    GetCards() ([]*Card, error) 
}

type CardUseCase interface {
    GetCards() ([]*Card,error)
}
