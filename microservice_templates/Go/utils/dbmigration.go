package main

import (
	"log"
	dbconfig "media-app-go/config"

	"github.com/joho/godotenv"
)

func main() {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	dbconfig.Connect()
	dbconfig.DB.AutoMigrate()
}
