package teacherHandler

import (
	teacherController "GolangAPIAssessment/internal/controller/teacher"
	"GolangAPIAssessment/internal/model"
	"GolangAPIAssessment/internal/utils"
	"net/http"

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
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": utils.ErrInternalServerError.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"teacher": newTeacher})
}
