package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nathanjcook/discordbotgo/bot"
	dbconfig "github.com/nathanjcook/discordbotgo/config"
)

func init() {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	// Connect to DB on app start up
	dbconfig.Connect()
}

func main() {

	bot.Start()

	<-make(chan struct{})
}
