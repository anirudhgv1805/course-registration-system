package controller

import (
	"backend-crs/model"
	"backend-crs/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	RegisterStudent(ctx *gin.Context)
}

type authController struct {
	studentService service.StudentService
}

func NewAuthController(s service.StudentService) AuthController {
	return &authController{studentService: s}
}


func (ctrl *authController) RegisterStudent(ctx *gin.Context){
	var student model.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error" : "Invalid Register Request",
		})
		return
	}

	registeredStudent, err := ctrl.studentService.RegisterStudent(&student)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable,gin.H{
			"error" : err.Error(),
			"status" : "failed",
		})
		panic("Error while registering the student "+err.Error())
	}

	ctx.JSON(http.StatusCreated,gin.H{
		"status" : "success",
		"body" : registeredStudent,
	})
}