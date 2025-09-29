package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Username     string     `json:"username" gorm:"unique"`
	Password     string     `json:"password"`
	RegisterNo   string     `json:"registerno" gorm:"unique"`
	DepartmentID uint       `json:"departmentId"`
	Department   Department `json:"department" gorm:"foreignKey:DepartmentID;references:ID"`
	Email        string     `json:"email"`
	Section      string     `json:"section"`
	Batch        int        `json:"batch"`
}
