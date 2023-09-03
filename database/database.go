package database

import (
	"GolangAPIAssessment/internal/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Singapore",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println("Successfully connected to the database")
	}

	migrations()
}

func migrations() {
	log.Println("Running Migrations")
	err := Database.AutoMigrate(&model.Student{}, &model.Teacher{})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println("Migrations completed")
	}
}
