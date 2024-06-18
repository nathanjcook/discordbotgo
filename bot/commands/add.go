package commands

import (
	"strconv"

	dbconfig "github.com/nathanjcook/discordbotgo/config"
	"go.uber.org/zap"
)

func Add(name string, url string, timeout string) string {
	type Microservice struct {
		MicroserviceId      int    `gorm:"column:microservice_id;unique;primaryKey;autoIncrement"`
		MicroserviceName    string `gorm:"column:microservice_name;size:25;"`
		MicroserviceUrl     string `gorm:"column:microservice_url;"`
		MicroserviceTimeout int    `gorm:"column:microservice_timeout;size:4;"`
	}
	var query Microservice
	var msg string

	timeout_int, err := strconv.Atoi(timeout)
	if err != nil {
		msg = "Timeout Is In An Incorrect Format"
		return msg
	} else {
		result := dbconfig.DB.Where("microservice_name = ? OR microservice_url = ?", name, url).Find(&query)
		if result.RowsAffected < 1 {
			microserviceAdd := Microservice{MicroserviceName: name, MicroserviceUrl: url, MicroserviceTimeout: timeout_int}
			err := dbconfig.DB.Create(&microserviceAdd).Error
			if err != nil {
				zap.L().Warn("[WARN]", zap.Error(err))
				msg = "Error Connecting To Database"
				return msg
			} else {
				msg = "Microservice: " + name + " Added To Server"
				return msg
			}
		} else {
			msg = "Microservice Microservice Name AND Microservice URL Must Be Unique"
			return msg
		}
	}
}
