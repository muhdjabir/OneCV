package teacherHandler

import (
	"GolangAPIAssessment/internal/model"
	"GolangAPIAssessment/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTeachers(ctx *gin.Context) {
	idparam := ctx.QueryArray("teacher")
	ctx.JSON(http.StatusOK, gin.H{"message": idparam})
}

func CreateTeacher(ctx *gin.Context) {
	var input model.TeacherInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": utils.ErrInvalidInputFormat.Error()})
		return
	}

	teacher := model.Teacher{
		Email: input.Email,
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": teacher})
}
