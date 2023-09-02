package main

import (
	"GolangAPIAssessment/database"
	"log"
)

func main() {
	log.Printf("This is my Golang log")
	database.Connect()
}
