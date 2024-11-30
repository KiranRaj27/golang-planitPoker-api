package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/kiranraj27/sprint-planner/config"
	_ "github.com/lib/pq"
)

func ConnectPostgres(cfg *config.Config) *sqlx.DB {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort,
	)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	return db
}
