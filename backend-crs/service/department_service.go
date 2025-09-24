package service

import (
	"backend-crs/dto"
	"backend-crs/mapper"
	"backend-crs/repository"
)

type DepartmentService interface {
	GetDepartments() (*[]dto.DepartmentResponse, error)
}

type departmentService struct {
	repo repository.DepartmentRepository
}

func NewDepartmentService(repo repository.DepartmentRepository) DepartmentService {
	return &departmentService{repo: repo}
}

func (s *departmentService) GetDepartments() (*[]dto.DepartmentResponse, error) {
	departments, err := s.repo.GetDepartments()
	if err != nil {
		return nil, err
	}
	var departmentResponses []dto.DepartmentResponse

	for _, department := range *departments {
		var departmentResponse dto.DepartmentResponse
		mapper.DepartmentToDepartmentResponseForListOfDepartments(&department, &departmentResponse)
		departmentResponses = append(departmentResponses, departmentResponse)
	}
	return &departmentResponses, nil
}
