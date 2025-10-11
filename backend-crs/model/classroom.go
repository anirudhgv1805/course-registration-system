package model

import "gorm.io/gorm"

type Classroom struct {
	gorm.Model
	Advisor1 Staff  `json:"advisor1"`
	Advisor2 Staff  `json:"advisor2"`
	RoomNo   string `json:"room_no"`
	Section  string `json:"section"`
}
