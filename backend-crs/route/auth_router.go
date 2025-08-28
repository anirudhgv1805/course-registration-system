package route

import (
	"backend-crs/controller"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(authController controller.AuthController, router *gin.Engine) {

	studentApi := router.Group("/api/student/auth")
	{
		studentApi.POST("/register", authController.RegisterStudent)
		studentApi.POST("/login", authController.LoginStudent)
	}
	staffApi := router.Group("/api/staff/auth")
	{
		staffApi.POST("/register", authController.RegisterStaff)
		staffApi.POST("/login", authController.LoginStaff)
	}
}
