package main

import (
	"github.com/gin-gonic/gin"

	"search-engine/config"
	"search-engine/routes"
)

func main() {
	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a new router instance
	router := gin.Default()

	// Set up the routes for the book search engine
	routes.SetupBookRoutes(router, db)

	// Start the server on port 8080
	router.Run(":6000")
}
