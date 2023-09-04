package studentController

import (
	"GolangAPIAssessment/database"
	"GolangAPIAssessment/internal/model"
	"GolangAPIAssessment/internal/utils"
)

func GetStudents(query []string) (model.StudentResponse, error) {
	var students []model.Student
	var studentResponse model.StudentResponse
	if len(query) != 0 {
		database.Database.Where("email IN ?", query).Find(&students)
		if len(students) != len(query) {
			return studentResponse, utils.ErrEntryNotFound
		}
	} else {
		database.Database.Find(&students)
	}
	for _, student := range students {
		studentResponse = append(studentResponse, student.Email)
	}
	return studentResponse, nil
}

func Create(student model.Student) (string, error) {
	if err := database.Database.Create(&student).Error; err != nil {
		return student.Email, utils.ErrInternalServerError
	}
	return student.Email, nil
}
