package addcommand

import (
	"strconv"

	dbconfig "github.com/nathanjcook/discordbotgo/config"
)

type Microservice struct {
	MicroserviceId      int    `gorm:"column:microservice_id;unique;primaryKey;autoIncrement"`
	MicroserviceName    string `gorm:"column:microservice_name;size:25;"`
	MicroserviceUrl     string `gorm:"column:microservice_url;"`
	MicroserviceTimeout int    `gorm:"column:microservice_timeout;size:4;"`
}

func Add(name string, url string, timeout string) (string, string) {
	var query Microservice
	var msg string
	var ttl string

	timeout_int, err := strconv.Atoi(timeout)
	if err != nil {
		ttl = "Add Command Error"
		msg = "Timeout Is In An Incorrect Format"
		return ttl, msg
	} else {
		result := dbconfig.DB.Where("microservice_name = ? OR microservice_url = ?", name, url).Find(&query)
		if result.RowsAffected < 1 {
			microserviceAdd := Microservice{MicroserviceName: name, MicroserviceUrl: url, MicroserviceTimeout: timeout_int}
			err := dbconfig.DB.Create(&microserviceAdd).Error
			if err != nil {
				ttl = "Add Command Error"
				msg = "Error Connecting To Database"
				return ttl, msg
			} else {
				ttl = "Add Command"
				msg = "Microservice: " + name + " Added To Server"
				return ttl, msg
			}
		} else {
			ttl = "Add Command Error"
			msg = "Microservice Name AND Microservice URL Must Be Unique"
			return ttl, msg
		}
	}
}
