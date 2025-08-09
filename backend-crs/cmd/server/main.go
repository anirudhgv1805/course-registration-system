package main

import (
	"backend-crs/config"
	"backend-crs/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config.LoadEnv()

	db, err := gorm.Open(postgres.Open(config.GetDBURI()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database " + err.Error())
	}

	db.AutoMigrate(&models.Student{}, &models.Admin{})

	router := gin.Default()
	
	router.GET("/me", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "The backend is working")
	})
	router.GET("/favicon.ico",func(ctx *gin.Context){
		ctx.File("./public/favicon.ico")
	})
	
	router.Run(":8080")
}
