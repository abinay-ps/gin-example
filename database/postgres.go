package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/abinay-ps/gin-example/config"
)

func Connect() *sql.DB {
	//Connect to the database
	db, err := sql.Open("postsgres", config.GetDbDetails())

	if err != nil {
		log.Fatal("Database connection failed")

	}

	if err := db.Ping(); err != nil {
		log.Fatal("Database Ping failed")
	}

	fmt.Printf("Database connection is successful")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS movies (id INTEGER PRIMARY KEY, title TEXT, director TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	//Return the database instance
	return db
}