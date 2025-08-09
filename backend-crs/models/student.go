package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	RegisterNo string `json:"registerno" gorm:"unique"`

}