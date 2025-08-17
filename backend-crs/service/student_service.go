package service

import (
	"backend-crs/model"
	"backend-crs/repository"
	"backend-crs/util"
	"errors"
)

var ErrorUserAlreadyExists = errors.New("The user trying to register is already registered!!")

type StudentService interface {
	RegisterStudent(student *model.Student) (*model.Student, error)
}

type studentService struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{repo: repo}
}

func (s *studentService) RegisterStudent(student *model.Student) (*model.Student, error) {
	if existing, _ := s.repo.FindStudentByRegisterNo(student.RegisterNo); existing.ID != 0 {
		return nil, ErrorUserAlreadyExists
	}

	hashedPassword, err := util.HashPassword(student.Password)
	if err != nil {
		panic("Failed while hashing the students password " + err.Error())
	}
	student.Password = hashedPassword

	err = s.repo.RegisterStudent(student)
	if err != nil {
		panic("Could not register the new student")
	}

	return student, err
}
