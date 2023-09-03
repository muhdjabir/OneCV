package main

import (
	"GolangAPIAssessment/database"
	"GolangAPIAssessment/serverd/router"
)

func main() {
	database.Connect()
	router.SetRoutes()
}
