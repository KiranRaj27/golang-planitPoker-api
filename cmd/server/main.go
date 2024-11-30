package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiranraj27/sprint-planner/config"
	"github.com/kiranraj27/sprint-planner/internal/users/https"
	user "github.com/kiranraj27/sprint-planner/internal/users/repository"
	"github.com/kiranraj27/sprint-planner/internal/users/usecase"
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
	appLogger.Info("Connected to PostgreSQL database")

	router := mux.NewRouter()

	userRepo := user.NewUserRepository(db)
	userUC := usecase.NewUserUseCase(userRepo)
	https.NewUserHandler(router, userUC)

	port := cfg.AppPort
	appLogger.Info(fmt.Sprintf("Server is running at port %s", port))
	log.Fatal(http.ListenAndServe(":"+port, router))
}
