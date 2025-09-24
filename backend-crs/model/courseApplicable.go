package model

import "gorm.io/gorm"

type CourseApplicable struct {
	gorm.Model
	CourseID     uint `gorm:"primaryKey"`
	DepartmentID uint `gorm:"primaryKey"`

	Course     Course     `gorm:"foreignKey:CourseID"`
	Department Department `gorm:"foreignKey:DepartmentID"`
}
