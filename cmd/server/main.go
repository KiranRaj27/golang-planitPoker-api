package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	cardHandler "github.com/kiranraj27/sprint-planner/cards/https"
	card "github.com/kiranraj27/sprint-planner/cards/repository"
	cardUseCase "github.com/kiranraj27/sprint-planner/cards/usecase"
	"github.com/kiranraj27/sprint-planner/config"
	userHandler "github.com/kiranraj27/sprint-planner/users/https"
	user "github.com/kiranraj27/sprint-planner/users/repository"
	userUseCase "github.com/kiranraj27/sprint-planner/users/usecase"

	"github.com/kiranraj27/sprint-planner/pkg/database"
	"github.com/kiranraj27/sprint-planner/pkg/logger"

	_ "github.com/kiranraj27/sprint-planner/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title PlanItPoker API
// @version 1.0
// @description API documentation for PlanItPoker clone.
// @host localhost:8080
// @BasePath /api
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

	
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	port := cfg.AppPort
	appLogger.Info(fmt.Sprintf("Server is running at port %s", port))
	log.Fatal(http.ListenAndServe(":"+port, router))
}
