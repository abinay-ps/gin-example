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

func GetDbDetails() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbName = os.Getenv("DB_NAME")

	fmt.Println("Environment variables loaded successfully")
	fmt.Println(dbHost, dbPort, dbUser, dbPass, dbName)
	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	return str
}

func GetPort() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	httpPort = os.Getenv("HTTP_PORT")
	if httpPort == "" {
		return ":8080" //Default Port
	}
	return ":" + httpPort
}
