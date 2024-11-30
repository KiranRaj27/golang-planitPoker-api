package config

import (
	"log"
	"os"
)

type Config struct {
    DBUser     string
    DBPassword string
    DBName     string
    DBHost     string
    DBPort     string
    AppPort    string
}

func LoadConfig() *Config {
    return &Config{
        DBUser:     getEnv("DB_USER", "admin"),
        DBPassword: getEnv("DB_PASSWORD", "password"),
        DBName:     getEnv("DB_NAME", "goclean"),
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        AppPort:    getEnv("APP_PORT", "8080"),
    }
}

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        log.Printf("Environment variable %s not set. Using default: %s", key, defaultValue)
        return defaultValue
    }
    return value
}
