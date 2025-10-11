package repository

import (
	"backend-crs/model"

	"gorm.io/gorm"
)

type ClassroomRepository interface {
	CreateClassroom(classroom *model.Classroom) (*model.Classroom, error)
}

type classroomRepository struct {
	db *gorm.DB
}

// func NewClassroomRepository(db *gorm.DB) ClassroomRepository {
// 	return &classroomRepository{db: db}
// }

