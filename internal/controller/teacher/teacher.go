package teacherController

import (
	"GolangAPIAssessment/database"
	"GolangAPIAssessment/internal/model"
)

func GetTeachers(query []string) (model.TeacherResponse, error) {
	var teachers []model.Teacher
	var teacherResponse model.TeacherResponse
	if len(query) != 0 {
		database.Database.Where("email IN ?", query).Find(&teachers)
	} else {
		database.Database.Find(&teachers)
	}
	for _, teacher := range teachers {
		teacherResponse = append(teacherResponse, teacher.Email)
	}
	return teacherResponse, nil
}

func Create(teacher model.Teacher) (string, error) {
	if err := database.Database.Create(&teacher).Error; err != nil {
		return teacher.Email, err
	}
	return teacher.Email, nil
}
