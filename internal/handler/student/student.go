package studentHandler

import (
	studentController "GolangAPIAssessment/internal/controller/student"
	"GolangAPIAssessment/internal/model"
	"GolangAPIAssessment/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStudents(ctx *gin.Context) {
	idparam := ctx.QueryArray("student")
	students, err := studentController.GetStudents(idparam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": utils.ErrInternalServerError.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"students": students})
}

func PostStudent(ctx *gin.Context) {
	var input model.StudentInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": utils.ErrInvalidInputFormat.Error()})
		return
	}

	student := model.Student{
		Email:     input.Email,
		Suspended: false,
	}
	newStudent, err := studentController.Create(student)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": utils.ErrInternalServerError.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"student": newStudent})
}
