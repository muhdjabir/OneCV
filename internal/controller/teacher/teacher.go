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
		if len(teachers) != len(query) {
			return teacherResponse, utils.ErrTeacherNotFound
		}
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
		return teacher.Email, utils.ErrDuplicateEntry
	}
	return teacher.Email, nil
}

func RegisterStudents(teacher string, students []string) error {
	var teacherRecord model.Teacher
	var studentList []model.Student

	database.Database.Where("email IN ?", students).Find(&studentList)
	err := database.Database.Where("email = ?", teacher).First(&teacherRecord).Error
	if err != nil {
		return utils.ErrTeacherNotFound
	}
	if len(studentList) != len(students) {
		return utils.ErrStudentNotFound
	}
	for _, student := range studentList {
		database.Database.Model(&student).Association("Teachers").Append(&teacherRecord)
	}
	return nil
}

func GetCommonStudents(teachers []string) (model.StudentResponse, error) {
	var commonStudents model.StudentResponse
	err := database.Database.
		Table("students").
		Select("students.email").
		Joins("INNER JOIN student_teachers ON students.id = student_teachers.student_id").
		Joins("INNER JOIN teachers ON student_teachers.teacher_id = teachers.id").
		Where("teachers.email IN ?", teachers).
		Group("students.email").
		Having("COUNT(DISTINCT teachers.email) = ?", len(teachers)).
		Pluck("students.email", &commonStudents).
		Error

	if err != nil {
		return nil, utils.ErrInternalServerError
	}

	return commonStudents, nil
}

func SuspendStudent(student string) error {
	var updatedStudent model.Student

	err := database.Database.Where("email = ?", student).First(&updatedStudent).Error
	if err != nil {
		return utils.ErrStudentNotFound
	}
	updatedStudent.Suspended = true
	database.Database.Save(updatedStudent)
	return nil
}

func RetrieveNotificationRecipients(teacher string, mentioned []string) (model.StudentResponse, error) {
	var students model.StudentResponse
	err := database.Database.
		Table("students").
		Select("DISTINCT students.email").
		Joins("LEFT JOIN student_teachers ON students.id = student_teachers.student_id").
		Where("students.suspended = false").
		Where("student_teachers.teacher_email = ?", teacher).
		Or("students.email IN ?", mentioned).
		Pluck("students.email", &students).
		Error

	if err != nil {
		return students, utils.ErrInternalServerError
	}
	return students, nil
}
