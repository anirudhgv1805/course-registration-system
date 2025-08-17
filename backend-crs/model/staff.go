package model

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	Username       string `json:"username"`
	Password       string `json:"password"`
	Department     string `json:"department"`
	Email          string `json:"email"`
	IsClassAdvisor bool   `json:"isClassAdvisor"`
	Section        string `json:"section"`
	Batch          string `json:"batch"`
}
