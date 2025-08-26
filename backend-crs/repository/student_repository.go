package repository

import (
	"backend-crs/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	RegisterStudent(student *model.Student) error
	FindStudentByRegisterNo(registerNo string) (*model.Student, error)
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db}
}

func (r *studentRepository) RegisterStudent(student *model.Student) error {
	return r.db.Create(student).Error
}

func (r *studentRepository) FindStudentByRegisterNo(registerNo string) (*model.Student, error) {
	var student model.Student
	err := r.db.Where("register_no = ?", registerNo).First(&student).Error
	return &student, err
}
