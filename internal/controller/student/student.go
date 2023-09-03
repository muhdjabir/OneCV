package studentController

import (
	"GolangAPIAssessment/database"
	"GolangAPIAssessment/internal/model"
)

func GetStudents(query []string) ([]model.Student, error) {
	var students []model.Student
	if len(query) != 0 {
		database.Database.Where("email IN ?", query).Find(&students)
	} else {
		database.Database.Find(&students)
	}
	return students, nil
}

func Create(student model.Student) (model.Student, error) {
	if err := database.Database.Create(&student).Error; err != nil {
		return student, err
	}
	return student, nil
}
