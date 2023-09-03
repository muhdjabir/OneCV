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

type TeacherActions struct {
	Teacher      string   `json:"teacher"`
	Student      string   `json:"student"`
	Students     []string `json:"students"`
	Notification string   `json:"notification"`
}
