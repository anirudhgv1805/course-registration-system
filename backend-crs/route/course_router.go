package route

import (
	"backend-crs/controller"
	"backend-crs/middleware"
	"backend-crs/ws"

	"github.com/gin-gonic/gin"
)

func SetupCourseRouter(courseController controller.CourseController, router *gin.Engine) {

	studentCourseApi := router.Group("/api/students/courses")
	studentCourseApi.Use(middleware.JWTAuthMiddleware())
	studentCourseApi.Use(middleware.AuthorizeRoles("student"))
	{
		studentCourseApi.GET("/", courseController.GetAllCourses)
		studentCourseApi.POST("/block", courseController.BlockSeat)
		studentCourseApi.POST("/apply", courseController.ApplyCourse)
	}
	staffCourseApi := router.Group("/api/staff/courses")
	staffCourseApi.Use(middleware.JWTAuthMiddleware())
	staffCourseApi.Use(middleware.AuthorizeRoles("staff"))
	{
		// course crud
		staffCourseApi.GET("/", courseController.GetAllCourses)
		staffCourseApi.POST("/", courseController.CreateCourse)
		staffCourseApi.PUT("/:id", courseController.UpdateCourseDetails)
		staffCourseApi.DELETE("/:id", courseController.DeleteCourse)

	}
}

func SetupRegisterRoutes(r *gin.Engine, cc controller.CourseController, hub *ws.Hub) {
	r.GET("/ws/course-updates", ws.ServeWS(hub))
	r.POST("/courses/:courseID/block", cc.BlockSeat)
}
