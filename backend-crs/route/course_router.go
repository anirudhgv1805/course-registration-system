package route

import (
	"backend-crs/controller"
	"backend-crs/middleware"

	"github.com/gin-gonic/gin"
)

func SetupCourseRouter(courseController controller.CourseController, router *gin.Engine) {

	studentCourseApi := router.Group("/api/students/courses")
	studentCourseApi.Use(middleware.JWTAuthMiddleware())
	{
		studentCourseApi.GET("/", courseController.GetAllCourses)
		studentCourseApi.POST("/block", courseController.BlockACourse)
		studentCourseApi.POST("/apply", courseController.ApplyCourse)
	}
	staffCourseApi := router.Group("/api/staff/courses")
	staffCourseApi.Use(middleware.JWTAuthMiddleware())

	{
		staffCourseApi.GET("/", courseController.GetAllCourses)
		staffCourseApi.POST("/",courseController.CreateCourse)
		staffCourseApi.PUT("/", courseController.UpdateCourseDetails)

	}
}
