package model

import "gorm.io/gorm"

type CourseApplicable struct {
	gorm.Model
	CourseID     uint
	DepartmentID uint

	Course     Course     `gorm:"foreignKey:CourseID"`
	Department Department `gorm:"foreignKey:DepartmentID"`
}
