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

	migrations(Database)
}

func SetUpMockDatabase() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Singapore",
		os.Getenv("TEST_DB_HOST"),
		os.Getenv("TEST_DB_USER"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"))
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println("Successfully connected to the database")
	}

	migrations(Database)
}

func TeardownMockDatabase() {
	Database.Migrator().DropTable(&model.Student{}, &model.Teacher{})
}

func migrations(database *gorm.DB) {
	log.Println("Running Migrations")
	err := database.AutoMigrate(&model.Student{}, &model.Teacher{})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println("Migrations completed")
	}
}
