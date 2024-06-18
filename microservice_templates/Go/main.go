package main

import (
	dbconfig "media-app-go/config"
	"media-app-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a gin router
	router := gin.Default()

	// Connect to DB on app start up
	dbconfig.Connect()

	// Provide access to routes
	routes.Routes(router)

	// Start the app
	router.Run(":8081")
}
