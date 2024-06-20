package commands

import (
	"fmt"
	"testing"

	dbconfig "github.com/nathanjcook/discordbotgo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MicroserviceInfo struct {
	MicroserviceId      int    `gorm:"column:microservice_id;unique;primaryKey;autoIncrement"`
	MicroserviceName    string `gorm:"column:microservice_name;size:25;"`
	MicroserviceUrl     string `gorm:"column:microservice_url;"`
	MicroserviceTimeout int    `gorm:"column:microservice_timeout;size:4;"`
}

func setupTestDBInfo() {
	host := "localhost"
	user := "postgres"
	password := "thorpe01685"
	dbname := "discord_db"
	port := "5433"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host,
		user,
		password,
		dbname,
		port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	dbconfig.DB = db

	db.AutoMigrate(&MicroserviceInfo{})
}

func TestInfoMS(t *testing.T) {
	setupTestDBInfo()

	dbconfig.DB.Create(&MicroserviceInfo{
		MicroserviceName:    "tester_test",
		MicroserviceUrl:     "http://localhost:3007",
		MicroserviceTimeout: 70,
	})

	title, msg := Info()
	title_dont_want := "Info Command Null"
	msg_dont_want := "No Microservices Available"

	if title_dont_want == title {
		t.Errorf("\n\nError: Info Failing To Get Microservice Data:\nWhat We Wanted: Info Command\nWhat We Got: %q", title)
	}
	if msg_dont_want == msg {
		t.Errorf("\n\nError: Info Failing To Get Microservice Data:\nWhat We Wanted: All Microservices\nWhat We Got: %q", msg)
	}
}
