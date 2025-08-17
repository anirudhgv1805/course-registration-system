// @title Course Registration System API
// @version 1.0
// @description This is the backend built with go for the course registration system

// @host      localhost:8080
// @BasePath  /

// @schemes http

package main

import (
	"backend-crs/config"
	"backend-crs/container"
	"net/http"

	"backend-crs/cmd/server/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

func main() {
	config.LoadEnv()


	router := gin.Default()

	router.GET("/me", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "The backend is working")
	})

	router.GET("/favicon.ico", func(ctx *gin.Context) {
		ctx.File("./public/favicon.ico")
	})

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	fx.New(container.Module).Run()

}
