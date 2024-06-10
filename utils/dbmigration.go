package main

import (
	"log"

	"github.com/joho/godotenv"
	dbconfig "github.com/nathanjcook/discordbotgo/config"
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
