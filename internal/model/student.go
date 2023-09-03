package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Email     string     `gorm:"primaryKey; not null; unique;" json:"email"`
	Suspended bool       `gorm:"not null;" json:"suspended"`
	Teachers  []*Teacher `gorm:"many2many:student_teachers;"`
}

type StudentInput struct {
	Email string `json:"email" binding:"required"`
}
