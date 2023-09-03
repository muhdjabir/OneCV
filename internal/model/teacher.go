package model

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Email string `gorm:"primaryKey; not null; unique;" json:"email"`
}

type TeacherInput struct {
	Email string `json:"email" binding:"required"`
}

type TeacherResponse = []string
