package conf

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConn() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres dbname=bookshop  password=admin port=5433  sslmode=disabled"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error in connecting to the database %v", err)
	}
	return db
}
