package controller

import (
	"backend-crs/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DepartmentController interface {
	GetDepartments(ctx *gin.Context)
}

type departmentController struct {
	departmentService service.DepartmentService
}

func NewDepartmentController(departmentService service.DepartmentService) DepartmentController {
	return &departmentController{departmentService: departmentService}
}

func (ctrl *departmentController) GetDepartments(ctx *gin.Context) {
	departments, err := ctrl.departmentService.GetDepartments()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "could not fetch departments",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":      "success",
		"departments": departments,
	})
}
