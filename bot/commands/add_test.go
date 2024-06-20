package commands

import (
	"fmt"
	"testing"

	dbconfig "github.com/nathanjcook/discordbotgo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MicroserviceAdd struct {
	MicroserviceId      int    `gorm:"column:microservice_id;unique;primaryKey;autoIncrement"`
	MicroserviceName    string `gorm:"column:microservice_name;size:25;"`
	MicroserviceUrl     string `gorm:"column:microservice_url;"`
	MicroserviceTimeout int    `gorm:"column:microservice_timeout;size:4;"`
}

func setupTestDBAdd() {
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

	db.AutoMigrate(&MicroserviceAdd{})
}

func TestAddMSNameAlreadyExists(t *testing.T) {
	setupTestDBAdd()

	dbconfig.DB.Create(&MicroserviceAdd{
		MicroserviceName:    "existing_service",
		MicroserviceUrl:     "http://localhost:3002",
		MicroserviceTimeout: 70,
	})

	title, msg := Add("existing_service", "http://localhost:8081", "50")
	title_want := "Add Command Error"
	msg_want := "Microservice Name AND Microservice URL Must Be Unique"

	if title_want != title {
		t.Errorf("\n\nError: Failed To Prevent User From Adding A Microservice Name That Already Exists:\nWhat We Wanted: %q\nWhat We Got: %q", title_want, title)
	}
	if msg_want != msg {
		t.Errorf("\n\nError: Failed To Prevent User From Adding A Microservice Name That Already Exists:\nWhat We Wanted: %q\nWhat We Got: %q", msg_want, msg)
	}
}

func TestAddMSHostURLAlreadyExists(t *testing.T) {
	setupTestDBAdd()

	dbconfig.DB.Create(&MicroserviceAdd{
		MicroserviceName:    "existing_service",
		MicroserviceUrl:     "http://localhost:3002",
		MicroserviceTimeout: 70,
	})

	title, msg := Add("New_service", "http://localhost:3002", "50")
	title_want := "Add Command Error"
	msg_want := "Microservice Name AND Microservice URL Must Be Unique"

	if title_want != title {
		t.Errorf("\n\nError: Failed To Prevent User From Adding A Microservice Host URL That Already Exists:\nWhat We Wanted: %q\nWhat We Got: %q", title_want, title)
	}
	if msg_want != msg {
		t.Errorf("\n\nError: Failed To Prevent User From Adding A Microservice Host URL That Already Exists:\nWhat We Wanted: %q\nWhat We Got: %q", msg_want, msg)
	}
}

func TestAddSuccess(t *testing.T) {
	setupTestDBAdd()

	dbconfig.DB.Create(&MicroserviceAdd{
		MicroserviceName:    "existing_service",
		MicroserviceUrl:     "http://localhost:3002",
		MicroserviceTimeout: 70,
	})

	title, msg := Add("New_service", "http://localhost:8081", "50")
	title_want := "Add Command"
	msg_want := "Microservice: New_service Added To Server"

	if title_want != title {
		t.Errorf("\n\nError: Failed To Add To Database Even If All Conditions Met:\nWhat We Wanted: %q\nWhat We Got: %q", title_want, title)
	} else if msg_want != msg {
		t.Errorf("\n\nError: Failed To Add To Database Even If All Conditions Met:\nWhat We Wanted: %q\nWhat We Got: %q", msg_want, msg)
	} else {
		Delete("New_service")
	}
}
