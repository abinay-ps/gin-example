package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	dbHost   string
	dbPort   string
	dbUser   string
	dbPass   string
	dbName   string
	httpPort string
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
	httpPort = os.Getenv("HTTP_PORT")
	fmt.Println("Environment variables loaded successfully")
}

func GetDbDetails() string {
	return "host=dbHost port=dbHost user=dbUser password=dbPass dbname=dbName sslmode=disable"

	//	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=12345 dbname=ScheduleManagement sslmode=disable")
}

func GetPort() string {
	if httpPort == "" {
		return ":8080" //Default Port
	}
	return ":" + httpPort
}
