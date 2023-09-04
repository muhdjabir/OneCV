package teacherHandler

import (
	teacherController "GolangAPIAssessment/internal/controller/teacher"
	"GolangAPIAssessment/internal/model"
	"GolangAPIAssessment/internal/utils"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetTeachers(ctx *gin.Context) {
	idparam := ctx.QueryArray("teacher")
	teachers, err := teacherController.GetTeachers(idparam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": utils.ErrInternalServerError.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello", "teachers": teachers})
}

func PostTeacher(ctx *gin.Context) {
	var input model.TeacherInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": utils.ErrInvalidInputFormat.Error()})
		return
	}

	teacher := model.Teacher{
		Email: input.Email,
	}
	newTeacher, err := teacherController.Create(teacher)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"teacher": newTeacher})
}

func RegisterStudents(ctx *gin.Context) {
	var input model.TeacherActions
	err := ctx.ShouldBindJSON(&input)
	if err != nil || len(input.Teacher) == 0 || len(input.Students) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": utils.ErrInvalidInputFormat.Error()})
		return
	}
	err = teacherController.RegisterStudents(input.Teacher, input.Students)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func GetCommonStudents(ctx *gin.Context) {
	idparam := ctx.QueryArray("teacher")
	if len(idparam) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": utils.ErrInvalidInputFormat})
		return
	}
	students, err := teacherController.GetCommonStudents(idparam)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"students": students})
}

func SuspendStudent(ctx *gin.Context) {
	var input model.TeacherActions
	err := ctx.ShouldBindJSON(&input)
	if err != nil || len(input.Student) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": utils.ErrInvalidInputFormat.Error()})
		return
	}
	err = teacherController.SuspendStudent(input.Student)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func RetrieveNotificationRecipients(ctx *gin.Context) {
	var input model.TeacherActions
	var emails []string
	err := ctx.ShouldBindJSON(&input)
	if err != nil || len(input.Teacher) == 0 || len(input.Notification) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": utils.ErrInvalidInputFormat.Error()})
		return
	}

	words := strings.Fields(input.Notification)
	for _, word := range words {
		if strings.HasPrefix(word, "@") {
			email := strings.TrimPrefix(word, "@")
			emails = append(emails, email)
		}
	}
	log.Println(emails)
	recipients, err := teacherController.RetrieveNotificationRecipients(input.Teacher, emails)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"recipients": recipients})
}
