package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
    err := godotenv.Load("../../.env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}

func GetEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
