package deletecommand

import (
	dbconfig "github.com/nathanjcook/discordbotgo/config"
)

type Microservice struct {
	MicroserviceId      int    `gorm:"column:microservice_id;unique;primaryKey;autoIncrement"`
	MicroserviceName    string `gorm:"column:microservice_name;size:25;"`
	MicroserviceUrl     string `gorm:"column:microservice_url;"`
	MicroserviceTimeout int    `gorm:"column:microservice_timeout;size:4;"`
}

func Delete(name string) (string, string) {
	var query Microservice
	var msg string
	var ttl string

	result := dbconfig.DB.Where("microservice_name = ?", name).Find(&query)
	if result.RowsAffected > 0 {
		dbconfig.DB.Where("microservice_name = ?", name).Delete(&Microservice{})
		ttl = "Delete Command"
		msg = "Microservice: " + name + " Has Been Deleted"
		return ttl, msg
	} else {
		ttl = "Delete Command Error"
		msg = "Bot Name Does Not Exist"
		return ttl, msg
	}
}
