package model

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	CourseCode      string `json:"courseCode"`
	Name            string `json:"name"`
	TotalCapacity   int    `json:"totalCapacity"`
	CurrentlyFilled int    `json:"currentlyFilled"`
	Department      string `json:"department"`
	StaffID         uint   `gorm:"uniqueIndex"`
	HandledByStaff  Staff  `json:"handledByStaff" gorm:"foreignKey:StaffID"`
	Semester        string `json:"semester"`
	Credit          int    `json:"credit"`
	Batch           int    `json:"batch"`
}
