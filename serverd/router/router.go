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

	public.GET("/commonstudents", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Common Students"})
	})
	public.POST("/register", teacherHandler.RegisterStudents)
	public.POST("/suspend", teacherHandler.SuspendStudent)
	public.POST("/retrievefornotifications", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Register Students"})
	})

	protected := router.Group("/api/admin")

	protected.GET("/student", studentHandler.GetStudents)
	protected.POST("/student", studentHandler.PostStudent)
	protected.GET("/teacher", teacherHandler.GetTeachers)
	protected.POST("/teacher", teacherHandler.PostTeacher)

	router.Run("0.0.0.0:8080")
}
