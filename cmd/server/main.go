package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiranraj27/sprint-planner/config"
	cardHandler "github.com/kiranraj27/sprint-planner/internal/cards/https"
	card "github.com/kiranraj27/sprint-planner/internal/cards/repository"
	cardUseCase "github.com/kiranraj27/sprint-planner/internal/cards/usecase"
	userHandler "github.com/kiranraj27/sprint-planner/internal/users/https"
	user "github.com/kiranraj27/sprint-planner/internal/users/repository"
	userUseCase "github.com/kiranraj27/sprint-planner/internal/users/usecase"

	"github.com/kiranraj27/sprint-planner/pkg/database"
	"github.com/kiranraj27/sprint-planner/pkg/logger"
)

func main() {
	cfg := config.LoadConfig()

	appLogger := logger.NewLogger()
	appLogger.Info("Starting the application...")

	db := database.ConnectPostgres(cfg)
	defer func() {
		if err := db.Close(); err != nil {
			appLogger.Error(fmt.Errorf("error closing database: %v", err))
		}
	}()
	err := database.SeedCards(db)
	if err != nil {
		log.Fatal("Error seeding cards: ", err)
	}
	appLogger.Info("Connected to PostgreSQL database")

	router := mux.NewRouter()

	userRepo := user.NewUserRepository(db)
	userUC := userUseCase.NewUserUseCase(userRepo)
	userHandler.NewUserHandler(router, userUC)

	cardRepo := card.NewCardRepository(db)
	cardUC := cardUseCase.NewCardUseCase(cardRepo)
	cardHandler.NewCardHandler(router, cardUC)

	

	port := cfg.AppPort
	appLogger.Info(fmt.Sprintf("Server is running at port %s", port))
	log.Fatal(http.ListenAndServe(":"+port, router))
}
