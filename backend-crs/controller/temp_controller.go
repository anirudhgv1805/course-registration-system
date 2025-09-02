package controller

import (
	"backend-crs/service"
	"backend-crs/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type TempController interface {
	VerifyToken(ctx *gin.Context)
}

type tempController struct {
	studentService service.StudentService
	staffService   service.StaffService
}

func NewTempController(studentService service.StudentService, staffService service.StaffService) TempController {
	return &tempController{studentService: studentService, staffService: staffService}
}

func (t *tempController) VerifyToken(ctx *gin.Context) {

	jwtToken := ctx.Request.Header.Get("Authorization")
	userData, err := util.ValidateJwtToken(strings.SplitAfter(jwtToken, "Bearer ")[1])
	if err != nil {
		ctx.JSON(http.StatusAccepted, gin.H{
			"userId": "",
			"error":  err,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"userId": userData["userId"],
		"role":   userData["role"],
	})
}
