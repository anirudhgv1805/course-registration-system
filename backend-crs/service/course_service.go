package service

import (
	"backend-crs/repository"
	"backend-crs/util"
)

type CourseService interface {
	UpdateCourseDetails(courseId int, updatedData map[string]interface{}) error
	GetAllCourseByDepartment()
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
