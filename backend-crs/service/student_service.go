package service

import (
	"backend-crs/dto"
	"backend-crs/model"
	"backend-crs/repository"
	"backend-crs/util"
	"errors"
)

var ErrorUserAlreadyExists = errors.New("the user trying to register is already registered")

type StudentService interface {
	RegisterStudent(student *model.Student) (*dto.StudentResponse, error)
	FindStudentByRegisterNo(registerNo string) (*model.Student, error)
	Authenticate(registerNo, password string) (*dto.StudentResponse, error)
}

type studentService struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{repo: repo}
}

func (s *studentService) RegisterStudent(student *model.Student) (*dto.StudentResponse, error) {
	if existing, _ := s.repo.FindStudentByRegisterNo(student.RegisterNo); existing.Username != "" {
		return nil, ErrorUserAlreadyExists
	}

	hashedPassword, err := util.HashPassword(student.Password)
	if err != nil {
		return nil, err
	}
	student.Password = hashedPassword

	err = s.repo.RegisterStudent(student)
	if err != nil {
		return nil, err
	}

	studentResponse := dto.StudentResponse{
		Username:   student.Username,
		RegisterNo: student.RegisterNo,
		Department: student.Department,
		Email:      student.Email,
		Section:    student.Section,
		Batch:      student.Batch,
	}
	return &studentResponse, err
}

func (s *studentService) FindStudentByRegisterNo(registerNo string) (*model.Student, error) {
	student, err := s.repo.FindStudentByRegisterNo(registerNo)
	if err != nil {
		return nil, err
	}
	return student, err
}

func (s *studentService) Authenticate(registerNo, password string) (*dto.StudentResponse, error) {
	student, err := s.repo.FindStudentByRegisterNo(registerNo)
	if err != nil {
		return nil, err
	}
	if !util.IsPasswordMatch(password, student.Password) {
		return nil, errors.New("invalid password")
	}

	return &dto.StudentResponse{
		Username:   student.Username,
		Email:      student.Email,
		Department: student.Department,
		Batch:      student.Batch,
		RegisterNo: student.RegisterNo,
		Section:    student.Section,
	}, nil
}
