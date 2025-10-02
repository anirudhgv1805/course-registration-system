package repository

import (
	"backend-crs/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	RegisterStudent(student *model.Student) (*model.Student, error)
	FindStudentByRegisterNo(registerNo string) (*model.Student, error)
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db}
}

func (r *studentRepository) RegisterStudent(student *model.Student) (*model.Student, error) {
	var savedStudent *model.Student
	if err := r.db.Create(savedStudent).Error; err != nil {
		return nil, err
	}
	if err := r.db.Preload("Department").Where("staff_id = ?", savedStudent.RegisterNo).Find(&savedStudent).Error; err != nil {
		return nil, err
	}
	return savedStudent, nil
}

func (r *studentRepository) FindStudentByRegisterNo(registerNo string) (*model.Student, error) {
	var student model.Student
	err := r.db.Preload("Department").Where("register_no = ?", registerNo).First(&student).Error
	return &student, err
}
