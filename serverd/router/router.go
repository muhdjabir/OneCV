package router

import (
	studentHandler "GolangAPIAssessment/internal/handler/student"
	teacherHandler "GolangAPIAssessment/internal/handler/teacher"
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

	protected.GET("/student", studentHandler.GetStudents)
	protected.POST("/student", studentHandler.CreateStudent)
	protected.GET("/teacher", teacherHandler.GetTeachers)
	protected.POST("/teacher", teacherHandler.CreateTeacher)

	router.Run("0.0.0.0:8080")
}
