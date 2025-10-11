package service

import (
	"backend-crs/dto"
	"backend-crs/mapper"
	"backend-crs/model"
	"backend-crs/redis"
	"backend-crs/repository"
	"backend-crs/util"
	"backend-crs/ws"
	"context"
	"errors"
)

type CourseService interface {
	UpdateCourseDetails(courseId int, updatedData map[string]interface{}) error
	GetAllCourseByDepartment()
	GetAllCourse() ([]dto.CourseResponse, error)
	CreateCourse(courseReq dto.CreateCourseDTO) (*model.Course, error)
	DeleteCourseById(courseId uint) error
	BlockCourse(registerNo string, courseId string) error
	BlockSeat(ctx context.Context, courseID, userID, deptID string) error
}

type courseService struct {
	repo  repository.CourseRepository
	Cache *redis.CourseCache
	Hub   *ws.Hub
}

func NewCourseService(repo repository.CourseRepository, cache *redis.CourseCache, hub *ws.Hub) CourseService {
	return &courseService{repo: repo, Cache: cache, Hub: hub}
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

func (c *courseService) BlockCourse(registerNo string, courseId string) error {
	panic("unimplemented")
}

func (cs *courseService) BlockSeat(ctx context.Context, courseID, userID, deptID string) error {
	ok, err := cs.Cache.BlockSeat(ctx, courseID, userID)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("seat already blocked")
	}

	// Example value - fetch actual seat count from DB
	updatedSeatCount := 10

	cs.Hub.BroadcastToDepartment(deptID, ws.CourseSeatUpdate{
		Type:      "seat_update",
		CourseID:  courseID,
		SeatsLeft: updatedSeatCount,
	})

	return nil
}
