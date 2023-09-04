package router

import (
	studentHandler "GolangAPIAssessment/internal/handler/student"
	teacherHandler "GolangAPIAssessment/internal/handler/teacher"
	"log"

	"github.com/gin-gonic/gin"
)

func SetRoutes() {
	router := gin.Default()

	public := router.Group("/api")

	public.GET("/commonstudents", teacherHandler.GetCommonStudents)
	public.POST("/register", teacherHandler.RegisterStudents)
	public.POST("/suspend", teacherHandler.SuspendStudent)
	public.POST("/retrievefornotifications", teacherHandler.RetrieveNotificationRecipients)

	protected := router.Group("/api/admin")

	protected.GET("/student", studentHandler.GetStudents)
	protected.POST("/student", studentHandler.PostStudent)
	protected.GET("/teacher", teacherHandler.GetTeachers)
	protected.POST("/teacher", teacherHandler.PostTeacher)

	router.Run("0.0.0.0:8080")
	log.Println("Server is listening at PORT:8080")
}
