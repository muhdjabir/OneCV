package main

import (
	"GolangAPIAssessment/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("This is my Golang log")
	database.Connect()
	setRoutes()
}

func setRoutes() {
	router := gin.Default()
	log.Println("Server is listening at PORT:8080")

	public := router.Group("/api")
	public.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to OneCV"})
	})

	router.Run("0.0.0.0:8080")
}
