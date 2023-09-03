package teacherController

import (
	"GolangAPIAssessment/database"
	"GolangAPIAssessment/internal/model"
	"GolangAPIAssessment/internal/utils"
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
		return teacher.Email, utils.ErrInternalServerError
	}
	return teacher.Email, nil
}

func RegisterStudents(teacher string, students []string) error {
	var teacherRecord model.Teacher
	var studentList []model.Student

	database.Database.Where("email IN ?", students).Find(&studentList)
	err := database.Database.Where("email = ?", teacher).First(&teacherRecord).Error
	if err != nil || len(studentList) != len(students) {
		return utils.ErrEntryNotFound
	}
	for _, student := range studentList {
		database.Database.Model(&student).Association("Teachers").Append(&teacherRecord)
	}
	return nil
}

func SuspendStudent(student string) error {
	var updatedStudent model.Student

	err := database.Database.Where("email = ?", student).First(&updatedStudent).Error
	if err != nil {
		return utils.ErrEntryNotFound
	}
	updatedStudent.Suspended = true
	database.Database.Save(updatedStudent)
	return nil
}
