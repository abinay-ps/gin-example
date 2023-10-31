package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	dbHost string
	dbPort string
	dbUser string
	dbPass string
	dbName string
)

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbName = os.Getenv("DB_NAME")
}

func GetDbDetails() string {
	return "host=dbHost port=dbHost user=dbUser password=dbPass dbname=dbName sslmode=disabled"
}
