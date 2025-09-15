package controller

import (
	"backend-crs/dto"
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
	BlockCourse(ctx *gin.Context)
	DeleteCourse(ctx *gin.Context)
}

type courseController struct {
	courseService service.CourseService
}

func NewCourseController(courseService service.CourseService) CourseController {
	return &courseController{courseService: courseService}
}

func (c *courseController) BlockCourse(ctx *gin.Context) {
	panic("unimplemented")
}

func (c *courseController) GetAllCourses(ctx *gin.Context) {
	courses, err := c.courseService.GetAllCourse()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"courses": courses,
	})
}

func (c *courseController) UpdateCourseDetails(ctx *gin.Context) {
	courseId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"status":  "failed",
			"message": err.Error(),
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
}

func (c *courseController) ApplyCourse(ctx *gin.Context) {
	panic("unimplemented")
}

func (c *courseController) GetAllCourseByDepartment(ctx *gin.Context) {
	panic("unimplemented")
}

func (c *courseController) CreateCourse(ctx *gin.Context) {
	var req dto.CreateCourseDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	course, err := c.courseService.CreateCourse(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "created course successfully",
		"course":  course,
	})
}

func (c *courseController) DeleteCourse(ctx *gin.Context) {
	courseId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}
	err = c.courseService.DeleteCourseById(uint(courseId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{
		"status":  "success",
		"message": "Deleted the course with ID " + strconv.FormatUint(courseId, 10),
	})
}
