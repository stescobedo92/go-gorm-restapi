package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DSN = "host=localhost user=stescobedo password=root dbname=gormdb port=5432"
)

var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB Connected")
	}
}
