package domain

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
