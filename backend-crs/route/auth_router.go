package route

import (
	"backend-crs/controller"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(authController controller.AuthController) *gin.Engine {

	router := gin.Default()

	api := router.Group("/api/student/auth")
	{
		api.POST("/register", authController.RegisterStudent)
	}
	return router
}
