package studentController

import (
	"GolangAPIAssessment/database"
	"GolangAPIAssessment/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStudents(t *testing.T) {
	database.SetUpMockDatabase()
	defer database.TeardownMockDatabase()

	// Populating the Test Database
	testStudent1 := model.Student{Email: "test1@gmail.com"}
	testStudent2 := model.Student{Email: "test2@gmail.com"}
	database.Database.Create(&testStudent1)
	database.Database.Create(&testStudent2)

	// Test case 1: Empty Query
	query := []string{}
	result, err := GetStudents(query)

	assert.NoError(t, err)
	assert.Contains(t, result, "test1@gmail.com")
	assert.Contains(t, result, "test2@gmail.com")

	// Test case 2: Get specified students
	query = []string{"test1@gmail.com"}
	result, err = GetStudents(query)

	assert.NoError(t, err)
	assert.Contains(t, result, "test1@gmail.com")
	assert.NotContains(t, result, "test2@gmail.com")

	// Test case 3: Students not found
	query = []string{"testnotfound@gmail.com"}
	_, err = GetStudents(query)

	assert.Error(t, err)
}

func TestCreateStudent(t *testing.T) {
	database.SetUpMockDatabase()
	defer database.TeardownMockDatabase()

	// Test case 1: Create user
	student1 := model.Student{Email: "test1@gmail.com"}
	result, err := Create(student1)

	assert.NoError(t, err)
	assert.Equal(t, "test1@gmail.com", result)

	// Test case 2: Error duplicate email
	student2 := model.Student{Email: "test1@gmail.com"}
	_, err = Create(student2)

	assert.Error(t, err)
}
