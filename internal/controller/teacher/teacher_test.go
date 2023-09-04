package teacherController

import (
	"GolangAPIAssessment/database"
	"GolangAPIAssessment/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTeachers(t *testing.T) {
	database.SetUpMockDatabase()
	defer database.TeardownMockDatabase()

	// Populating the Test Database
	testTeacher1 := model.Teacher{Email: "test1@gmail.com"}
	testTeacher2 := model.Teacher{Email: "test2@gmail.com"}
	database.Database.Create(&testTeacher1)
	database.Database.Create(&testTeacher2)

	// Test case 1: Empty Query
	query := []string{}
	result, err := GetTeachers(query)

	assert.NoError(t, err)
	assert.Contains(t, result, "test1@gmail.com")
	assert.Contains(t, result, "test2@gmail.com")

	// Test case 2: Get specified teachers
	query = []string{"test1@gmail.com"}
	result, err = GetTeachers(query)

	assert.NoError(t, err)
	assert.Contains(t, result, "test1@gmail.com")
	assert.NotContains(t, result, "test2@gmail.com")

	// Test case 3: Teachers not found
	query = []string{"testnotfound@gmail.com"}
	_, err = GetTeachers(query)

	assert.Error(t, err)
}

func TestCreateTeacher(t *testing.T) {
	database.SetUpMockDatabase()
	defer database.TeardownMockDatabase()

	// Test case 1: Create user
	teacher1 := model.Teacher{Email: "test1@gmail.com"}
	result, err := Create(teacher1)

	assert.NoError(t, err)
	assert.Equal(t, "test1@gmail.com", result)

	// Test case 2: Error duplicate email
	teacher2 := model.Teacher{Email: "test1@gmail.com"}
	_, err = Create(teacher2)

	assert.Error(t, err)
}

func TestRegisterStudents(t *testing.T) {
	database.SetUpMockDatabase()
	defer database.TeardownMockDatabase()

	// Populating database
	teacher1 := model.Teacher{Email: "teacherken@gmail.com"}
	student1 := model.Student{Email: "studentjon@gmail.com"}
	student2 := model.Student{Email: "studenthon@gmail.com"}
	database.Database.Create(&teacher1)
	database.Database.Create(&student1)
	database.Database.Create(&student2)

	// Test case 1: Successful register of students
	err := RegisterStudents(teacher1.Email, []string{student1.Email, student2.Email})
	assert.NoError(t, err)

	// // Test case 2: Teacher not registered
	err = RegisterStudents("teacherunregistered@gmail.com", []string{student1.Email, student2.Email})
	assert.Error(t, err)

	// // Test case 3: Unregistered student
	err = RegisterStudents(teacher1.Email, []string{student1.Email, "studentunregistered@gmail.com"})
	assert.Error(t, err)
}

func TestGetCommonStudents(t *testing.T) {
	database.SetUpMockDatabase()
	defer database.TeardownMockDatabase()

	// Populating database
	teacher1 := model.Teacher{Email: "teacherken@gmail.com"}
	student1 := model.Student{Email: "studentjon@gmail.com"}
	student2 := model.Student{Email: "studenthon@gmail.com"}
	database.Database.Create(&teacher1)
	database.Database.Create(&student1)
	database.Database.Create(&student2)
	database.Database.Model(&student1).Association("Teachers").Append(&teacher1)
	database.Database.Model(&student2).Association("Teachers").Append(&teacher1)

	// Test case 1: Get students with registered teacher
	result, err := GetCommonStudents([]string{teacher1.Email})

	assert.NoError(t, err)
	assert.Contains(t, result, student1.Email)
	assert.Contains(t, result, student2.Email)

	// Test case 2: Unregistered teacher or no students under teacher
	result, err = GetCommonStudents([]string{"teacherunregistered@gmail.com"})

	assert.NoError(t, err)
	assert.Equal(t, len(result), 0)
}

func TestSuspendStudent(t *testing.T) {
	database.SetUpMockDatabase()
	defer database.TeardownMockDatabase()

	// Populating database
	student1 := model.Student{Email: "studentjon@gmail.com"}
	student2 := model.Student{Email: "studenthon@gmail.com"}
	database.Database.Create(&student1)

	// Test case 1: Successfully suspend a student
	err := SuspendStudent(student1.Email)

	assert.NoError(t, err)

	// Test case 2: Failure to suspend an unregistered student
	err = SuspendStudent(student2.Email)

	assert.Error(t, err)
}

func TestRetrieveNotificationRecipients(t *testing.T) {
	database.SetUpMockDatabase()
	defer database.TeardownMockDatabase()

	// Populating database
	teacher1 := model.Teacher{Email: "teacherken@gmail.com"}
	student1 := model.Student{Email: "studentjon@gmail.com"}
	student2 := model.Student{Email: "studenthon@gmail.com"}
	database.Database.Create(&teacher1)
	database.Database.Create(&student1)
	database.Database.Create(&student2)
	database.Database.Model(&student1).Association("Teachers").Append(&teacher1)

	// Test case 1: Only students under the teacher
	result, err := RetrieveNotificationRecipients(teacher1.Email, []string{})

	assert.NoError(t, err)
	assert.Contains(t, result, "studentjon@gmail.com")
	assert.NotContains(t, result, "studenthon@gmail.com")

	// Test case 2: Student under teacher is suspended
	student1.Suspended = true
	database.Database.Save(student1)
	result, err = RetrieveNotificationRecipients(teacher1.Email, []string{})

	assert.NoError(t, err)
	assert.NotContains(t, result, "studentjon@gmail.com")
	assert.NotContains(t, result, "studenthon@gmail.com")

	// Test case 3: Student is mentioned
	result, err = RetrieveNotificationRecipients(teacher1.Email, []string{student2.Email})

	assert.NoError(t, err)
	assert.NotContains(t, result, "studentjon@gmail.com")
	assert.Contains(t, result, "studenthon@gmail.com")
}
