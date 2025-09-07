package controller

import (
	"backend-crs/dto"
	"backend-crs/model"
	"backend-crs/service"
	"backend-crs/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	RegisterStudent(ctx *gin.Context)
	RegisterStaff(ctx *gin.Context)
	LoginStudent(ctx *gin.Context)
	LoginStaff(ctx *gin.Context)
}

type authController struct {
	studentService service.StudentService
	staffService   service.StaffService
}

func NewAuthController(studentService service.StudentService, staffService service.StaffService) AuthController {
	return &authController{studentService: studentService, staffService: staffService}
}

func (ctrl *authController) RegisterStudent(ctx *gin.Context) {
	var student model.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Invalid Register Request",
		})
		return
	}

	registeredStudent, err := ctrl.studentService.RegisterStudent(&student)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"error":  err.Error(),
			"status": "failed",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"body":   registeredStudent,
	})
}

func (ctrl *authController) LoginStudent(ctx *gin.Context) {
	var loginData dto.LoginDTO
	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Invalid Login Request",
		})
		return
	}

	studentResponse, err := ctrl.studentService.Authenticate(loginData.UserId, loginData.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}
	jwtToken, err := util.GenerateJwtToken(loginData.UserId, "student")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"status": "success",
		"token":  jwtToken,
		"body":   studentResponse,
	})
}

// staff related methods

func (ctrl *authController) RegisterStaff(ctx *gin.Context) {
	var staff model.Staff
	if err := ctx.ShouldBindJSON(&staff); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	registeredStaff, err := ctrl.staffService.RegisterStaff(&staff)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"error":  err.Error(),
			"status": "failed",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"body":   registeredStaff,
	})
}

func (ctrl *authController) LoginStaff(ctx *gin.Context) {
	var loginData dto.LoginDTO
	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "Invalid Register Request",
		})
		return
	}

	studentResponse, err := ctrl.staffService.Authenticate(loginData.UserId, loginData.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}
	jwtToken, err := util.GenerateJwtToken(loginData.UserId, "staff")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"status": "success",
		"token":  jwtToken,
		"body":   studentResponse,
	})
}
