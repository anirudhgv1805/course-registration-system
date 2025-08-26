package repository

import (
	"backend-crs/model"

	"gorm.io/gorm"
)

type StaffRepository interface {
	RegisterStaff(staff *model.Staff) error
	FindStaffByStaffId(staffId string) (*model.Staff, error)
}

type staffRepository struct {
	db *gorm.DB
}

func (r *staffRepository) RegisterStaff(staff *model.Staff) error {
	return r.db.Create(staff).Error
}

func (r *staffRepository) FindStaffByStaffId(staffId string) (*model.Staff, error) {
	var staff model.Staff
	err := r.db.Where("staff_id = ?",staffId).First(&staff).Error
	return &staff, err
}