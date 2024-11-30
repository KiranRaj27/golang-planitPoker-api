package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func SeedCards(db *sqlx.DB) error {
	queries := []string{
		`INSERT INTO cards (value, set_type) VALUES
			('1', 'scrum'),
			('2', 'scrum'),
			('3', 'scrum'),
			('5', 'scrum'),
			('8', 'scrum'),
			('13', 'scrum'),
			('21', 'scrum'),
			('34', 'scrum'),
			('55', 'scrum'),
			('89', 'scrum')`,

		`INSERT INTO cards (value, set_type) VALUES
			('1', 'fibonacci'),
			('2', 'fibonacci'),
			('3', 'fibonacci'),
			('5', 'fibonacci'),
			('8', 'fibonacci'),
			('13', 'fibonacci'),
			('21', 'fibonacci'),
			('34', 'fibonacci'),
			('55', 'fibonacci'),
			('89', 'fibonacci')`,

		`INSERT INTO cards (value, set_type) VALUES
			('0.5', 'hours'),
			('1', 'hours'),
			('2', 'hours'),
			('3', 'hours'),
			('5', 'hours'),
			('8', 'hours'),
			('13', 'hours'),
			('20', 'hours'),
			('40', 'hours')`,

		`INSERT INTO cards (value, set_type) VALUES
			('XS', 'tshirt'),
			('S', 'tshirt'),
			('M', 'tshirt'),
			('L', 'tshirt'),
			('XL', 'tshirt'),
			('XXL', 'tshirt')`,

		`INSERT INTO cards (value, set_type) VALUES
			('1', 'sequential'),
			('2', 'sequential'),
			('3', 'sequential'),
			('4', 'sequential'),
			('5', 'sequential'),
			('6', 'sequential'),
			('7', 'sequential'),
			('8', 'sequential'),
			('9', 'sequential')`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Printf("Error seeding data: %v", err)
			return err
		}
	}
	log.Println("Database seeded successfully!")
	return nil
}
