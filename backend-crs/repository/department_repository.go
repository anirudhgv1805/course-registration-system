package repository

import (
	"backend-crs/model"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	GetDepartments() (*[]model.Department, error)
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{db}
}

func (r *departmentRepository) GetDepartments() (*[]model.Department, error) {
	var departments []model.Department
	err := r.db.Find(&departments).Error
	return &departments, err

}
