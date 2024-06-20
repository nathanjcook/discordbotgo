package commands

import (
	dbconfig "github.com/nathanjcook/discordbotgo/config"
)

func Info() (string, string) {
	type Microservice struct {
		MicroserviceId      int    `gorm:"column:microservice_id;unique;primaryKey;autoIncrement"`
		MicroserviceName    string `gorm:"column:microservice_name;size:25;"`
		MicroserviceUrl     string `gorm:"column:microservice_url;"`
		MicroserviceTimeout int    `gorm:"column:microservice_timeout;size:4;"`
	}
	var names []string
	msg := ""
	dbconfig.DB.Model(&Microservice{}).Pluck("microservice_name", &names)
	if len(names) > 0 {
		title := "Info Command"
		for i := 0; i < len(names); i++ {
			msg += "!gobot " + names[i] + " help\n\n"
		}
		return title, msg
	} else {
		title := "Info Command Null"
		return title, "No Microservices Available"
	}
}
