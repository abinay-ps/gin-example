package main

import (
	"github.com/abinay-ps/gin-example/config"
	"github.com/abinay-ps/gin-example/database"
	"github.com/abinay-ps/gin-example/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	//Load all configarations
	config.Load()
	//connect to the database
	db := database.Connect()

	defer db.Close()
	//create a router for handlers
	router := gin.Default()

	//Initialize routes
	handlers.Iniatilize_routes(router, db)

}
