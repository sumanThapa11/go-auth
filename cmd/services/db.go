package services

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConnection() *gorm.DB {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
