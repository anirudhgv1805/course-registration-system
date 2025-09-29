package model

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	Username       string `json:"username"`
	StaffId        string `json:"staffId"`
	Password       string `json:"password"`
	DepartmentID   uint   `json:"departmentId"`
	Department     Department `json:"department" gorm:"foreignKey:DepartmentID"`
	Email          string `json:"email"`
	IsClassAdvisor bool   `json:"isClassAdvisor"`
	Section        string `json:"section"`
	Batch          int    `json:"batch"`
}
