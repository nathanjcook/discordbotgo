package dbconfig

import (
	"gorm.io/gorm"
)

// Create var of type grom.DB
var DB *gorm.DB

func Connect() {
	// Get env variable
	// host := os.Getenv("POSTGRES_HOST")
	// user := os.Getenv("POSTGRES_USER")
	// password := os.Getenv("POSTGRES_PASSWORD")
	// dbname := os.Getenv("DATABASE_NAME")
	// port := os.Getenv("POSTGRES_PORT")

	// // data source name for postgres db
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
	// 	host,
	// 	user,
	// 	password,
	// 	dbname,
	// 	port)
	// // Connect to postgres with gorm
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	DB = db
	// }
}
