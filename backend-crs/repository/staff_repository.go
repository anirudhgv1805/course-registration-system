package repository

import (
	"backend-crs/model"

	"gorm.io/gorm"
)

type StaffRepository interface {
	RegisterStaff(staff *model.Staff) (*model.Staff, error)
	FindStaffByStaffId(staffId string) (*model.Staff, error)
}

type staffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) StaffRepository {
	return &staffRepository{db}
}

func (r *staffRepository) RegisterStaff(staff *model.Staff) (*model.Staff, error) {
	var savedStaff *model.Staff
	if err := r.db.Create(staff).Error; err != nil {
		return nil, err
	}
	if err := r.db.Preload("Department").Where("staff_id = ?", staff.StaffId).Find(&savedStaff).Error; err != nil {
		return nil, err
	}
	return savedStaff, nil
}

func (r *staffRepository) FindStaffByStaffId(staffId string) (*model.Staff, error) {
	var staff model.Staff
	err := r.db.Where("staff_id = ?", staffId).First(&staff).Error
	return &staff, err
}
