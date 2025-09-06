package model

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name  string `gorm:"unique" json:"name"`
	Code  string `json:"code"`
	Block string `json:"block"`
}
