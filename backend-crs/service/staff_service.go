package service

import (
	"backend-crs/dto"
	"backend-crs/mapper"
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

func NewStaffService(repo repository.StaffRepository) StaffService {
	return &staffService{repo: repo}
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

	staffResponse := &dto.StaffResponse{
		Username:       staff.Username,
		Email:          staff.Email,
		IsClassAdvisor: staff.IsClassAdvisor,
		Section:        staff.Section,
		Batch:          staff.Batch,
		Role:           "staff",
		StaffId:        staff.StaffId,
	}

	mapper.DepartmentToDepartmentResponseForListOfDepartments(&staff.Department, &staffResponse.Department)

	return staffResponse, nil
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

	savedStaff, err := s.repo.RegisterStaff(staff)
	if err != nil {
		return nil, err
	}

	staffResponse := dto.StaffResponse{
		Username:       savedStaff.Username,
		StaffId:        savedStaff.StaffId,
		Email:          savedStaff.Email,
		Section:        savedStaff.Section,
		Batch:          savedStaff.Batch,
		IsClassAdvisor: savedStaff.IsClassAdvisor,
		Role:           "staff",
	}
	mapper.DepartmentToDepartmentResponse(&savedStaff.Department, &staffResponse.Department)
	return &staffResponse, nil
}
