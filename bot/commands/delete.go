package commands

import (
	dbconfig "github.com/nathanjcook/discordbotgo/config"
)

func Delete(name string) string {
	type Microservice struct {
		MicroserviceId      int    `gorm:"column:microservice_id;unique;primaryKey;autoIncrement"`
		MicroserviceName    string `gorm:"column:microservice_name;size:25;"`
		MicroserviceUrl     string `gorm:"column:microservice_url;"`
		MicroserviceTimeout int    `gorm:"column:microservice_timeout;size:4;"`
	}
	var query Microservice
	var msg string

	result := dbconfig.DB.Where("microservice_name = ?", name).Find(&query)
	if result.RowsAffected > 0 {
		dbconfig.DB.Where("microservice_name = ?", name).Delete(&Microservice{})
		msg = "Microservice: " + name + " Has Been Deleted"
		return msg
	} else {
		msg = "Bot Name Does Not Exist"
		return msg
	}
}
