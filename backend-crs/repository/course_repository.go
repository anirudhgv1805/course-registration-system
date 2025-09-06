package repository

import (
	"backend-crs/model"

	"gorm.io/gorm"
)

type CourseRepository interface {
	CreateCourse(course *model.Course) error
	UpdateCourseDetails(courseId int, updatedData map[string]interface{}) error
	DeleteCourse(course *model.Course) error
	FindCourseByCourseCode(courseId string) (*model.Course, error)
	FindCourseByDepartment(department string) (*model.Course, error)
	GetAllCourse() ([]model.Course, error)
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db}
}
func (c *courseRepository) CreateCourse(course *model.Course) error {
	return c.db.Create(course).Error
}

func (c *courseRepository) DeleteCourse(course *model.Course) error {
	return c.db.Delete(course).Error
}

func (c *courseRepository) FindCourseByCourseCode(courseCode string) (*model.Course, error) {
	var course model.Course
	err := c.db.Where("course_code = ?", courseCode).First(course).Error
	return &course, err
}

func (c *courseRepository) FindCourseByDepartment(department string) (*model.Course, error) {
	var course model.Course
	err := c.db.Where("course_code = ?", department).First(&course).Error
	return &course, err
}

func (c *courseRepository) UpdateCourseDetails(courseId int, updatedData map[string]interface{}) error {
	var course model.Course
	err := c.db.Where("ID = ?", courseId).First(&course).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *courseRepository) GetAllCourse() ([]model.Course, error) {
	var courses []model.Course
	if err := c.db.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}
