package service

import (
	"backend-crs/dto"
	"backend-crs/model"
	"backend-crs/repository"
	"backend-crs/util"
	"errors"
)

type StaffService interface {
	FindStaffByStaffId(staffId string) (*model.Staff, error)
	Authenticate(staffId, password string) (*dto.StaffResponse, error)
	RegisterStaff(staff *model.Staff) (*dto.StaffResponse, error)
}

type staffService struct {
	repo repository.StaffRepository
}

func (s *staffService) FindStaffByStaffId(staffId string) (*model.Staff, error) {
	staff, err := s.repo.FindStaffByStaffId(staffId)
	if err != nil {
		return nil, err
	}
	return staff, nil

}

func (s *staffService) Authenticate(staffId, password string) (*dto.StaffResponse, error) {
	staff, err := s.repo.FindStaffByStaffId(staffId)
	if err != nil {
		return nil, err
	}
	if !util.IsPasswordMatch(password, staff.Password) {
		return nil, errors.New("invalid password")
	}

	return &dto.StaffResponse{
		Username:       staff.Username,
		Department:     staff.Department,
		Email:          staff.Email,
		IsClassAdvisor: staff.IsClassAdvisor,
		Section:        staff.Section,
		Batch:          staff.Batch,
	}, nil
}

func (s *staffService) RegisterStaff(staff *model.Staff) (*dto.StaffResponse, error) {
	if existing, _ := s.repo.FindStaffByStaffId(staff.StaffId); existing.Username != "" {
		return nil, ErrorUserAlreadyExists
	}

	hashedPassword, err := util.HashPassword(staff.Password)
	if err != nil {
		return nil, err
	}
	staff.Password = hashedPassword

	err = s.repo.RegisterStaff(staff)
	if err != nil {
		return nil, err
	}

	staffResponse := dto.StaffResponse{
		Username: staff.Username,
		StaffId : staff.StaffId,
		Department: staff.Department,
		Email: staff.Email,
		Section: staff.Section,
		Batch: staff.Batch,
		IsClassAdvisor: staff.IsClassAdvisor,
	}

	return &staffResponse, nil
}
