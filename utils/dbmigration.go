package main

import (
	"github.com/joho/godotenv"
	dbconfig "github.com/nathanjcook/discordbotgo/config"
	"go.uber.org/zap"
)

type Microservice struct {
	MicroserviceId      int    `gorm:"column:microservice_id;unique;primaryKey;autoIncrement"`
	MicroserviceName    string `gorm:"column:microservice_name;size:25;"`
	MicroserviceUrl     string `gorm:"column:microservice_url;"`
	MicroserviceTimeout int    `gorm:"column:microservice_timeout;size:4;"`
}

func main() {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		zap.L().Fatal("Error loading .env file:", zap.Error(err))
	}

	dbconfig.Connect()
	dbconfig.DB.AutoMigrate(&Microservice{})
}
