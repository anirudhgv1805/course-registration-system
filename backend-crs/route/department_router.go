package route

import (
	"backend-crs/controller"

	"github.com/gin-gonic/gin"
)

func SetupDepartmentRouter(departmentController controller.DepartmentController, router *gin.Engine) {
	departmentApi := router.Group("/api")
	{
		departmentApi.GET("/departments", departmentController.GetDepartments)  // matches /api/departments

	}
}
