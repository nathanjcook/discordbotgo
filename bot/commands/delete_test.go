package commands

import (
	"fmt"
	"testing"

	dbconfig "github.com/nathanjcook/discordbotgo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MicroserviceDelete struct {
	MicroserviceId      int    `gorm:"column:microservice_id;unique;primaryKey;autoIncrement"`
	MicroserviceName    string `gorm:"column:microservice_name;size:25;"`
	MicroserviceUrl     string `gorm:"column:microservice_url;"`
	MicroserviceTimeout int    `gorm:"column:microservice_timeout;size:4;"`
}

func setupTestDBDelete() {
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

	db.AutoMigrate(&MicroserviceDelete{})
}

func TestDeleteMockExist(t *testing.T) {
	setupTestDBDelete()

	dbconfig.DB.Create(&MicroserviceDelete{
		MicroserviceName: "existing_service",
	})

	title, msg := Delete("existing_service")
	title_want := "Delete Command"
	msg_want := "Microservice: existing_service Has Been Deleted"

	if title_want != title {
		t.Errorf("\n\nError: Delete Failed For Existing Microservice:\nWhat We Wanted: %q\nWhat We Got: %q", title_want, title)
	}
	if msg_want != msg {
		t.Errorf("\n\nError: Delete Failed For Existing Microservice:\nWhat We Wanted: %q\nWhat We Got: %q", msg_want, msg)
	}
}

func TestDeleteBadInput(t *testing.T) {
	setupTestDBDelete()

	title, msg := Delete("adsadadsadssadsaddsa")
	title_want := "Delete Command Error"
	msg_want := "Bot Name Does Not Exist"

	if title_want != title {
		t.Errorf("\n\nError: Bot still trying to delete non existent bot\nWhat We Wanted: %q\nWhat We Got: %q", title_want, title)
	}
	if msg_want != msg {
		t.Errorf("\n\nError: Bot still trying to delete non existent bot\nWhat We Wanted: %q\nWhat We Got: %q", msg_want, msg)
	}
}
