package main

import (
	"github.com/nathanjcook/discordbotgo/bot"
)

func init() {
	// Find .env file
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %s", err)
	// }

	// Connect to DB on app start up
	// dbconfig.Connect()
}

func main() {

	bot.Start()

	<-make(chan struct{})
}
