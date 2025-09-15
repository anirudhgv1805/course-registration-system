package service

import (
	"backend-crs/dto"
	"backend-crs/mapper"
	"backend-crs/model"
	"backend-crs/repository"
	"backend-crs/util"
)

type CourseService interface {
	UpdateCourseDetails(courseId int, updatedData map[string]interface{}) error
	GetAllCourseByDepartment()
	GetAllCourse() ([]dto.CourseResponse, error)
	CreateCourse(courseReq dto.CreateCourseDTO) (*model.Course, error)
	DeleteCourseById(courseId uint) error
}

type courseService struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{repo: repo}
}

func (c *courseService) UpdateCourseDetails(courseId int, updatedData map[string]any) error {
	newData := make(map[string]any)

	allowedCourseFields := map[string]any{
		"courseCode":      "string",
		"name":            "string",
		"totalCapacity":   "int",
		"currentlyFilled": "int",
		"department":      "string",
		"staffId":         "uint",
		"semester":        "string",
		"credit":          "int",
		"batch":           "int",
	}

	validatedData, err := util.FilterAllowedFieldsWithTypes(allowedCourseFields, newData)
	if err != nil {
		return err
	}

	err = c.repo.UpdateCourseDetails(courseId, validatedData)
	if err != nil {
		return err
	}

	return nil
}

func (c *courseService) GetAllCourseByDepartment() {
	panic("unimplemented")
}

func (c *courseService) GetAllCourse() ([]dto.CourseResponse, error) {
	courses, err := c.repo.GetAllCourses()
	if err != nil {
		return nil, err
	}
	courseResponses := make([]dto.CourseResponse, 0, len(courses))
	for _, course := range courses {
		var courseResponse dto.CourseResponse
		mapper.CourseToCourseResponseDTO(&course, &courseResponse)
		courseResponses = append(courseResponses, courseResponse)
	}
	return courseResponses, nil
}

func (c *courseService) CreateCourse(courseReq dto.CreateCourseDTO) (*model.Course, error) {

	course := model.Course{
		CourseCode:      courseReq.CourseCode,
		Name:            courseReq.Name,
		TotalCapacity:   courseReq.TotalCapacity,
		CurrentlyFilled: 0,
		DepartmentID:    courseReq.DepartmentID,
		StaffID:         courseReq.StaffID,
		Credit:          courseReq.Credit,
	}
	if err := c.repo.CreateCourse(&course); err != nil {
		return nil, err
	}

	return &course, nil
}

func (c *courseService) DeleteCourseById(courseId uint) error {
	var course model.Course
	course.ID = courseId
	if err := c.repo.DeleteCourse(&course); err != nil {
		return err
	}
	return nil
}
