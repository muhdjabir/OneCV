package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRoutes() {
	router := gin.Default()
	log.Println("Server is listening at PORT:8080")

	public := router.Group("/api")

	public.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to OneCV"})
	})
	public.GET("/register", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Register Students"})
	})
	public.GET("/commonstudents", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Common Students"})
	})
	public.GET("/suspend", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Suspend Students"})
	})
	public.GET("/retrievefornotifications", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Register Students"})
	})

	protected := router.Group("/api/admin")

	protected.GET("/student", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Get Students"})
	})
	protected.POST("/student", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Create new student"})
	})
	protected.GET("/teacher", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Get Teachers"})
	})
	protected.POST("/teacher", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Create new teacher"})
	})

	router.Run("0.0.0.0:8080")
}
