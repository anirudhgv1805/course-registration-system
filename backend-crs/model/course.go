package model

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	CourseCode            string             `json:"courseCode"`
	Name                  string             `json:"name"`
	TotalCapacity         int                `json:"totalCapacity"`
	CurrentlyFilled       int                `json:"currentlyFilled"`
	DepartmentID          uint               `json:"departmentId"`
	OfferedBy             Department         `json:"department" gorm:"foreignKey:DepartmentID"`
	StaffID               uint               `gorm:"uniqueIndex"`
	HandledByStaff        Staff              `json:"handledByStaff" gorm:"foreignKey:StaffID"`
	Semester              int                `json:"semester"`
	Credit                int                `json:"credit"`
	Batch                 int                `json:"batch"`
	ApplicableDepartments []CourseApplicable `gorm:"courseID"`
}
