package controller

import (
	"backend-crs/model"
	"backend-crs/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CourseController interface {
	GetAllCourses(ctx *gin.Context)            //This method is to handle the requests from admin and staff requests
	GetAllCourseByDepartment(ctx *gin.Context) // This method is to handle requests from students
	CreateCourse(ctx *gin.Context)
	UpdateCourseDetails(ctx *gin.Context)
	ApplyCourse(ctx *gin.Context)
	BlockACourse(ctx *gin.Context)
}

type courseController struct {
	courseService service.CourseService
}

func NewCourseController(courseService service.CourseService) CourseController {
	return &courseController{courseService: courseService}
}

func (c *courseController) BlockACourse(ctx *gin.Context) {
	panic("unimplemented")
}

func (c *courseController) GetAllCourses(ctx *gin.Context) {
	panic("unimplemented")
}

func (c *courseController) UpdateCourseDetails(ctx *gin.Context) {
	courseId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"status":  "failed",
			"message": "invalid id passed as param",
		})
		return
	}
	newData := make(map[string]any)

	err = ctx.ShouldBindJSON(&newData)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"status":  "failed",
			"message": "invalid course update",
		})
		return
	}
	err = c.courseService.UpdateCourseDetails(int(courseId), newData)
	if err != nil {
		ctx.JSON(http.StatusAccepted, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "course details updated successfully",
	})
	return
}

func (c *courseController) ApplyCourse(ctx *gin.Context) {
	panic("unimplemented")
}

func (c *courseController) GetAllCourseByDepartment(ctx *gin.Context) {
	panic("unimplemented")
}

func (c *courseController) CreateCourse(ctx *gin.Context) {
	var course model.Course

	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":"failed",
			"message":"course creation request was given with unknown values",
		})
		return
	}

	



	return
}
